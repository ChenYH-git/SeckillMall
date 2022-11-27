package controller

import (
	"context"
	"cyh_project/cyh_product_srv/data_source"
	"cyh_project/cyh_product_srv/models"
	cyh_product_srv "cyh_project/cyh_product_srv/proto/product"
	"errors"
	"time"
)

type Products struct{}

func (p *Products) ProductList(ctx context.Context, in *cyh_product_srv.ProductsRequest, out *cyh_product_srv.ProductsResponse) error {
	currentPage, pageSize := in.CurrentPage, in.PageSize
	var products []models.Products
	var product []models.Products
	res := data_source.Db.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&products)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "获取商品列表失败"
		return res.Error
	}

	total := int32(data_source.Db.Find(&product).RowsAffected)
	var protoProductResp []*cyh_product_srv.Product
	for _, p := range products {
		productResp := &cyh_product_srv.Product{
			Id:         int32(p.Id),
			Price:      p.Price,
			Num:        int32(p.Num),
			Name:       p.Name,
			Unit:       p.Unit,
			Pic:        p.Pic,
			Desc:       p.Desc,
			CreateTime: p.CreateTime.Format("2006-01-02 15:04:05"),
		}
		protoProductResp = append(protoProductResp, productResp)
	}
	out.Code = 200
	out.Msg = "获取商品列表成功"
	out.Products = protoProductResp
	out.CurrentPage = currentPage
	out.PageSize = pageSize
	out.Total = total
	return nil
}

func (p *Products) ProductAdd(ctx context.Context, in *cyh_product_srv.ProductAddRequest, out *cyh_product_srv.ProductAddResponse) error {
	name, price := in.Name, in.Price
	num, unit := in.Num, in.Unit
	picPath, desc := in.Pic, in.Desc

	product := models.Products{
		Num:        int(num),
		Price:      price,
		Name:       name,
		Unit:       unit,
		Pic:        picPath,
		Desc:       desc,
		CreateTime: time.Now(),
	}
	res := data_source.Db.Create(&product)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "添加商品失败"
		return nil
	}
	out.Code = 200
	out.Msg = "添加商品成功"
	return nil
}

func (p *Products) ProductDel(ctx context.Context, in *cyh_product_srv.ProductDelRequest, out *cyh_product_srv.ProductDelResponse) error {
	id := in.Id
	var products models.Products
	res := data_source.Db.Where("id = ?", int(id)).Delete(&products)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "删除商品失败"
		return nil
	}
	if res.RowsAffected < 1 {
		out.Code = 500
		out.Msg = "不存在此商品"
		return nil
	}
	out.Code = 200
	out.Msg = "删除商品成功"
	return nil
}

func (p *Products) ProductToEdit(ctx context.Context, in *cyh_product_srv.ProductToEditRequest, out *cyh_product_srv.ProductToEditResponse) error {
	id := in.Id

	product := models.Products{Id: int(id)}
	res := data_source.Db.First(&product)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "未查询到对应商品"
		return errors.New(out.Msg)
	}

	productResp := &cyh_product_srv.Product{
		Id:    int32(product.Id),
		Num:   int32(product.Num),
		Price: product.Price,
		Name:  product.Name,
		Unit:  product.Unit,
		Pic:   product.Pic,
		Desc:  product.Desc,
	}
	out.Code = 200
	out.Msg = "编辑商品成功"
	out.Product = productResp
	return nil
}

func (p *Products) ProductDoEdit(ctx context.Context, in *cyh_product_srv.ProductEditRequest, out *cyh_product_srv.ProductEditResponse) error {
	name, price, id := in.Name, in.Price, in.Id
	num, unit := in.Num, in.Unit
	picPath, desc := in.Pic, in.Desc

	var product, tmp models.Products
	product = models.Products{
		Num:   int(num),
		Price: price,
		Name:  name,
		Unit:  unit,
		Desc:  desc,
	}
	if len(picPath) >= 1 {
		product.Pic = picPath
	}
	res := data_source.Db.Find(&tmp).Where("id = ?", int(id)).Update(&product)
	if res.Error != nil {
		out.Code = 500
		out.Msg = "修改商品失败"
		return nil
	}
	out.Code = 200
	out.Msg = "修改商品成功"
	return nil
}
