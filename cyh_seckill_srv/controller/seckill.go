package controller

import (
	"cyh_project/cyh_seckill_srv/data_source"
	"cyh_project/cyh_seckill_srv/models"
	"cyh_project/cyh_seckill_srv/redis_lib"
	"cyh_project/cyh_web/utils"
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

//type SecKill struct{}

//func (s *SecKill) FrontSecKill(ctx context.Context, in *cyh_seckill_srv.SecKillRequest, out *cyh_seckill_srv.SecKillResponse) error {
//	id, curTime := in.Id, time.Now()
//	userEmail := in.Username
//	var seckill models.SecKills
//	res := data_source.Db.Where("id = ? AND start_time <= ? AND end_time >= ?", int(id), curTime, curTime).Find(&seckill)
//	if res.Error != nil {
//		out.Code = 500
//		out.Msg = "查询对应活动商品失败"
//		return nil
//	}
//	var order models.Orders
//
//	findRes := data_source.Db.Where("s_id = ? AND uemail = ?", int(id), userEmail).Find(&order)
//	if findRes.RowsAffected > 0 {
//		out.Code = 500
//		out.Msg = "同一用户只能购买一次"
//		return nil
//	}
//
//	order.SId, order.Uemail = int(id), userEmail
//	order.CreateTime = time.Now()
//	orderRes := data_source.Db.Create(&order)
//	if orderRes.Error != nil {
//		out.Code = 500
//		out.Msg = "创建订单失败"
//		return nil
//	}
//
//	upRes := res.Where("num > ?", 0).Update("num", seckill.Num-1)
//	if upRes.Error != nil {
//		out.Code = 500
//		out.Msg = "商品已售罄"
//		return nil
//	}
//	out.Code = 200
//	out.Msg = "下单成功"
//	return nil
//}

func InitConsumerMq() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()

	ch.Qos(1, 0, false)

	deliveries, err := ch.Consume("cyh_web.order_queue", "order_consumer", false, false, false, false, nil)
	if err != nil {
		fmt.Println(err)
	}

	for delivery := range deliveries {
		go OrderApply(delivery)
	}
}

func OrderApply(delivery amqp.Delivery) {
	body := string(delivery.Body)
	requestData := utils.StrToMap(body)

	id := requestData["pid"].(string)
	frontUserEmail := requestData["uemail"].(string) + id

	var seckill models.SecKills
	currTime := time.Now()
	res := data_source.Db.Where("id = ? AND start_time <= ? AND end_time >= ?", int(utils.StrToInt32(id)), currTime, currTime).Find(&seckill)
	if res.Error != nil || res.RowsAffected < 1 {
		delivery.Ack(true)
		mapData := map[string]interface{}{
			"code": 500,
			"msg":  "不存在此活动",
		}
		redis_lib.Conn.Do("SET", frontUserEmail, utils.MapToStr(mapData))
		return
	}

	var orderMq models.Orders
	if data_source.Db.Where("s_id = ? AND uemail = ?", int(utils.StrToInt32(id)), frontUserEmail).Find(&orderMq).RowsAffected > 0 {
		delivery.Ack(true)
		mapData := map[string]interface{}{
			"code": 500,
			"msg":  "同一用户只可购买一次",
		}
		redis_lib.Conn.Do("SET", frontUserEmail, utils.MapToStr(mapData))
		return
	}

	orderMq.SId, orderMq.Uemail = int(utils.StrToInt32(id)), frontUserEmail
	orderMq.CreateTime = time.Now()
	orderRes := data_source.Db.Create(&orderMq)
	if orderRes.Error != nil {
		delivery.Ack(true)
		mapData := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败，请重试",
		}
		redis_lib.Conn.Do("SET", frontUserEmail, utils.MapToStr(mapData))
		return
	}

	upRes := res.Where("num > ?", 0).Update("num", seckill.Num-1)
	if upRes.Error != nil {
		delivery.Ack(true)
		mapData := map[string]interface{}{
			"code": 500,
			"msg":  "下单失败，请重试",
		}
		redis_lib.Conn.Do("SET", frontUserEmail, utils.MapToStr(mapData))
		return
	}

	delivery.Ack(true)
	mapData := map[string]interface{}{
		"code": 200,
		"msg":  "下单成功",
	}
	redis_lib.Conn.Do("SET", frontUserEmail, utils.MapToStr(mapData))
	return
}
