package seckill

import (
	"context"
	cyh_product_srv "cyh_project/cyh_product_srv/proto/seckill"
	"cyh_project/cyh_web/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetSecKillList(ctx *gin.Context) {
	currentPage, pageSize := ctx.DefaultQuery("currentPage", "1"), ctx.DefaultQuery("pageSize", "5")
	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, err := seckillService.SecKillList(context.TODO(), &cyh_product_srv.SecKillsRequest{
		CurrentPage: utils.StrToInt32(currentPage),
		PageSize:    utils.StrToInt32(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         resp.Code,
		"msg":          resp.Msg,
		"seckills":     resp.Seckills,
		"total":        resp.Total,
		"current_page": resp.CurrentPage,
		"page_size":    resp.PageSize,
	})
}

func GetProducts(ctx *gin.Context) {

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, err := seckillService.GetProducts(context.TODO(), &cyh_product_srv.ProductRequest{})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":     resp.Code,
		"msg":      resp.Msg,
		"products": resp.Products,
	})
}

func AddSecKill(ctx *gin.Context) {
	name, price := ctx.PostForm("name"), ctx.PostForm("price")
	num, pid := ctx.PostForm("num"), ctx.PostForm("pid")
	startTime, endTime := ctx.PostForm("start_time"), ctx.PostForm("end_time")

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, _ := seckillService.SecKillAdd(context.TODO(), &cyh_product_srv.SecKill{
		Num:       utils.StrToInt32(num),
		Pid:       utils.StrToInt32(pid),
		Price:     utils.StrToFloat32(price),
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func DelSecKill(ctx *gin.Context) {
	id := ctx.PostForm("id")

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, _ := seckillService.SecKillDel(context.TODO(), &cyh_product_srv.SecKillDelRequest{Id: utils.StrToInt32(id)})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func ToEditSecKill(ctx *gin.Context) {
	id := ctx.Query("id")

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, err := seckillService.SecKillToEdit(context.TODO(), &cyh_product_srv.SecKillDelRequest{Id: utils.StrToInt32(id)})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":        resp.Code,
		"msg":         resp.Msg,
		"seckill":     resp.Seckill,
		"products_no": resp.ProductsNo,
	})
}

func EditSecKill(ctx *gin.Context) {
	name, price := ctx.PostForm("name"), ctx.PostForm("price")
	num, id := ctx.PostForm("num"), ctx.PostForm("id")
	pid := ctx.PostForm("pid")
	startTime, endTime := ctx.PostForm("start_time"), ctx.PostForm("end_time")

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())

	resp, _ := seckillService.SecKillDoEdit(context.TODO(), &cyh_product_srv.SecKill{
		Id:        utils.StrToInt32(id),
		Pid:       utils.StrToInt32(pid),
		Num:       utils.StrToInt32(num),
		Price:     utils.StrToFloat32(price),
		Name:      name,
		StartTime: startTime,
		EndTime:   endTime,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func GetFrontSeckillList(ctx *gin.Context) {
	currentPage := ctx.DefaultQuery("currentPage", "1")
	pageSize := ctx.DefaultQuery("pageSize", "8")

	service := micro.NewService()
	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, err := seckillService.FrontSecKillList(context.TODO(), &cyh_product_srv.FrontSecKillRequest{
		CurrentPage: utils.StrToInt32(currentPage),
		Pagesize:    utils.StrToInt32(pageSize),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	for _, seckill := range resp.Seckills {
		seckill.Pic = utils.ImgToBase64(seckill.Pic)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":         resp.Code,
		"msg":          resp.Msg,
		"current":      resp.CurrentPage,
		"page_size":    resp.PageSize,
		"total_page":   resp.Total,
		"seckill_list": resp.Seckills,
	})
}

func SecKillDetail(ctx *gin.Context) {
	id := ctx.Query("id")

	service := micro.NewService()

	seckillService := cyh_product_srv.NewSecKillsService("cyh_product_srv", service.Client())
	resp, err := seckillService.FrontSecKillDetail(context.TODO(), &cyh_product_srv.SecKillDelRequest{
		Id: utils.StrToInt32(id),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	resp.Seckill.Pic = utils.ImgToBase64(resp.Seckill.Pic)
	ctx.JSON(http.StatusOK, gin.H{
		"code":    resp.Code,
		"msg":     resp.Msg,
		"seckill": resp.Seckill,
	})
}
