package controller

import (
	"context"
	"cyh_project/cyh_product_srv/data_source"
	"cyh_project/cyh_product_srv/models"
	cyh_product_srv "cyh_project/cyh_product_srv/proto/seckill"
	"cyh_project/cyh_product_srv/utils"
	"errors"
	"time"
)

type SecKills struct{}

func (s *SecKills) SecKillList(ctx context.Context, in *cyh_product_srv.SecKillsRequest, out *cyh_product_srv.SecKillsResponse) error {
	currentPage, pageSize := in.CurrentPage, in.PageSize
	var seckills []models.SecKills
	var seckill []models.SecKills
	res := data_source.Db.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&seckills)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "获取活动列表失败"
		return res.Error
	}

	total := int32(data_source.Db.Find(&seckill).RowsAffected)
	var protoSecKillResp []*cyh_product_srv.SecKill
	for _, s := range seckills {
		product := models.Products{Id: s.PId}
		data_source.Db.First(&product)
		seckillResp := &cyh_product_srv.SecKill{
			Id:         int32(s.Id),
			Pid:        int32(s.PId),
			Price:      s.Price,
			Num:        int32(s.Num),
			Name:       s.Name,
			Pname:      product.Name,
			StartTime:  s.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    s.EndTime.Format("2006-01-02 15:04:05"),
			CreateTime: s.CreateTime.Format("2006-01-02 15:04:05"),
		}
		protoSecKillResp = append(protoSecKillResp, seckillResp)
	}
	out.Code = 200
	out.Msg = "获取活动列表成功"
	out.Seckills = protoSecKillResp
	out.CurrentPage = currentPage
	out.PageSize = pageSize
	out.Total = total
	return nil
}

func (s *SecKills) GetProducts(ctx context.Context, in *cyh_product_srv.ProductRequest, out *cyh_product_srv.ProductResponse) error {
	var products []models.Products

	res := data_source.Db.Find(&products)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "查询可选商品失败"
		return res.Error
	}

	var productsProtoResp []*cyh_product_srv.ProductSelected
	for _, product := range products {
		productProtoResp := &cyh_product_srv.ProductSelected{
			Id:    int32(product.Id),
			Pname: product.Name,
		}
		productsProtoResp = append(productsProtoResp, productProtoResp)
	}

	out.Code = 200
	out.Msg = "查询可选商品成功"
	out.Products = productsProtoResp
	return nil
}

func (s *SecKills) SecKillAdd(ctx context.Context, in *cyh_product_srv.SecKill, out *cyh_product_srv.SecKillResponse) error {
	name, price := in.Name, in.Price
	num, pid := in.Num, in.Pid
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)

	seckills := models.SecKills{
		PId:        int(pid),
		Num:        int(num),
		Price:      price,
		Name:       name,
		StartTime:  startTime,
		EndTime:    endTime,
		CreateTime: time.Now(),
	}
	res := data_source.Db.Create(&seckills)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "添加活动失败"
		return nil
	}
	out.Code = 200
	out.Msg = "添加活动成功"
	return nil
}

func (s *SecKills) SecKillDel(ctx context.Context, in *cyh_product_srv.SecKillDelRequest, out *cyh_product_srv.SecKillResponse) error {
	id := in.Id

	var seckill models.SecKills

	res := data_source.Db.Where("id = ?", int(id)).Delete(&seckill)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "删除活动失败"
		return nil
	}

	out.Code = 200
	out.Msg = "删除活动成功"
	return nil
}

func (s *SecKills) SecKillToEdit(ctx context.Context, in *cyh_product_srv.SecKillDelRequest, out *cyh_product_srv.SecKillToEditResponse) error {
	id := in.Id

	seckill := models.SecKills{Id: int(id)}
	res := data_source.Db.First(&seckill)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "未查询到对应活动"
		return errors.New(out.Msg)
	}

	var product models.Products
	data_source.Db.Where("id = ?", seckill.PId).Find(&product)
	productResp := &cyh_product_srv.SecKill{
		Id:        int32(seckill.Id),
		Pid:       int32(seckill.PId),
		Num:       int32(seckill.Num),
		Price:     seckill.Price,
		Name:      seckill.Name,
		Pname:     product.Name,
		StartTime: seckill.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:   seckill.EndTime.Format("2006-01-02 15:04:05"),
	}
	var productNoResp []*cyh_product_srv.ProductSelected
	var productsResp []models.Products
	data_source.Db.Find(&productsResp)
	for _, productNo := range productsResp {
		product := &cyh_product_srv.ProductSelected{
			Id:    int32(productNo.Id),
			Pname: productNo.Name,
		}
		productNoResp = append(productNoResp, product)
	}

	out.Code = 200
	out.Msg = "编辑商品成功"
	out.Seckill = productResp
	out.ProductsNo = productNoResp
	return nil
}

func (s *SecKills) SecKillDoEdit(ctx context.Context, in *cyh_product_srv.SecKill, out *cyh_product_srv.SecKillResponse) error {
	name, price := in.Name, in.Price
	num, id := in.Num, in.Id
	pid := in.Pid
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)

	var seckills, tmp models.SecKills
	seckills = models.SecKills{
		Num:       int(num),
		Price:     price,
		Name:      name,
		PId:       int(pid),
		StartTime: startTime,
		EndTime:   endTime,
	}
	res := data_source.Db.Find(&tmp).Where("id = ?", int(id)).Update(&seckills)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "修改活动失败"
		return nil
	}
	out.Code = 200
	out.Msg = "修改活动成功"
	return nil
}

func (s *SecKills) FrontSecKillList(ctx context.Context, in *cyh_product_srv.FrontSecKillRequest, out *cyh_product_srv.SecKillsResponse) error {
	tomorrowTime := utils.AddHour(24)

	var seckills []models.SecKills
	result := data_source.Db.Where("start_time <= ?", tomorrowTime).Where("status = ?", 0).Limit(in.Pagesize).Offset(in.Pagesize * (in.CurrentPage - 1)).Find(&seckills)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "查询不到活动数据"
		return errors.New(out.Msg)
	}

	var seckillsResp []*cyh_product_srv.SecKill
	for _, seckill := range seckills {
		var product models.Products
		data_source.Db.Where("id = ?", seckill.PId).Find(&product)
		seckillResp := &cyh_product_srv.SecKill{
			Id:         int32(seckill.Id),
			Num:        int32(seckill.Num),
			Pid:        int32(seckill.PId),
			Price:      seckill.Price,
			Pname:      product.Name,
			Name:       seckill.Name,
			Pic:        product.Pic,
			PPrice:     product.Price,
			Pdesc:      product.Desc,
			StartTime:  seckill.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    seckill.EndTime.Format("2006-01-02 15:04:05"),
			CreateTime: seckill.CreateTime.Format("2006-01-02 15:04:05"),
		}
		seckillsResp = append(seckillsResp, seckillResp)
	}

	var seckillsCount []models.SecKills
	count := int32(data_source.Db.Where("start_time <= ?", tomorrowTime).Where("status = ?", 0).Find(&seckillsCount).RowsAffected)

	out.Code = 200
	out.Msg = "查询活动数据成功"
	out.CurrentPage = in.CurrentPage
	out.PageSize = in.Pagesize
	out.Total = (count + in.Pagesize - 1) / in.Pagesize
	out.Seckills = seckillsResp
	return nil
}

func (s *SecKills) FrontSecKillDetail(ctx context.Context, in *cyh_product_srv.SecKillDelRequest, out *cyh_product_srv.FrongSecKillDetailResponse) error {
	id := in.Id

	var seckill models.SecKills
	result := data_source.Db.Where("id = ?", id).Find(&seckill)
	if result.Error != nil {
		out.Code = 500
		out.Msg = "没有查询到具体活动数据"
		return errors.New(out.Msg)
	}

	var product models.Products
	data_source.Db.Where("id = ?", seckill.PId).Find(&product)
	seckillResp := &cyh_product_srv.SecKill{
		Id:         int32(seckill.Id),
		Num:        int32(seckill.Num),
		Pid:        int32(seckill.PId),
		Price:      seckill.Price,
		Pname:      product.Name,
		Name:       seckill.Name,
		Pic:        product.Pic,
		PPrice:     product.Price,
		Pdesc:      product.Desc,
		StartTime:  seckill.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:    seckill.EndTime.Format("2006-01-02 15:04:05"),
		CreateTime: seckill.CreateTime.Format("2006-01-02 15:04:05"),
	}

	out.Code = 200
	out.Msg = "查询具体活动数据成功"
	out.Seckill = seckillResp
	return nil
}
