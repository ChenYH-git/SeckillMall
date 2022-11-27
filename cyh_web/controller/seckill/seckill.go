package seckill

import (
	"cyh_project/cyh_web/rabbitmq"
	"cyh_project/cyh_web/utils"
	"fmt"
	"net/http"

	"github.com/gomodule/redigo/redis"

	"github.com/gin-gonic/gin"
)

func SecKill(ctx *gin.Context) {
	id := ctx.PostForm("id")
	email, ok := ctx.Get("frontUserEmail")
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "获取用户信息失败，请登录",
		})
		return
	}

	qe := &rabbitmq.QueueAndExchange{
		QueueName:    "cyh_web.order_queue",
		ExchangeName: "cyh_web.order_exchange",
		ExchangeType: "direct",
		RoutingKey:   "cyh_web.order",
	}

	mq := rabbitmq.NewRabbitMq(qe)

	mq.ConnMq()
	mq.OpenChan()
	defer mq.CloseMq()
	defer mq.CloseChan()

	orderMap := map[string]interface{}{
		"uemail": email,
		"pid":    id,
	}

	mq.PublishMsg(utils.MapToStr(orderMap))

	ctx.JSON(http.StatusOK, gin.H{
		"code": 500,
		"msg":  "下单中，请稍后",
	})
}

func GetSeckillResult(ctx *gin.Context) {
	uemail, exist := ctx.Get("frontUserEmail")
	id := ctx.Query("id")
	if !exist {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "用户未登录",
		})
		return
	}

	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("连接出错")
		return
	}
	defer conn.Close()

	ret, err := redis.String(conn.Do("get", uemail.(string)+id))
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
		})
		return
	}
	retMap := utils.StrToMap(ret)
	ctx.JSON(http.StatusOK, gin.H{ // 说明从redis里面获取到了数据，
		"code": 200,
		"msg":  retMap["msg"],
	})
	return
}
