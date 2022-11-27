package product

import (
	"context"
	cyh_product_srv "cyh_project/cyh_product_srv/proto/product"
	"cyh_project/cyh_web/utils"
	"net/http"
	"strconv"
	"time"

	"go-micro.dev/v4/logger"

	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
)

func GetProductList(ctx *gin.Context) {
	currentPage, pageSize := ctx.DefaultQuery("currentPage", "1"), ctx.DefaultQuery("pageSize", "5")
	service := micro.NewService()
	productService := cyh_product_srv.NewProductsService("cyh_product_srv", service.Client())
	resp, err := productService.ProductList(context.TODO(), &cyh_product_srv.ProductsRequest{
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
		"products":     resp.Products,
		"total":        resp.Total,
		"current_page": resp.CurrentPage,
		"page_size":    resp.PageSize,
	})
}

func AddProduct(ctx *gin.Context) {
	name, price := ctx.PostForm("name"), ctx.PostForm("price")
	num, unit := ctx.PostForm("num"), ctx.PostForm("unit")
	desc := ctx.PostForm("desc")

	service := micro.NewService()
	productService := cyh_product_srv.NewProductsService("cyh_product_srv", service.Client())

	file, err := ctx.FormFile("pic")
	if err != nil {
		logger.Error(err)
		resp, _ := productService.ProductAdd(context.TODO(), &cyh_product_srv.ProductAddRequest{
			Num:   utils.StrToInt32(num),
			Price: utils.StrToFloat32(price),
			Name:  name,
			Unit:  unit,
			Desc:  desc,
		})

		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	filePath := "upload/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	ctx.SaveUploadedFile(file, filePath)

	resp, _ := productService.ProductAdd(context.TODO(), &cyh_product_srv.ProductAddRequest{
		Num:   utils.StrToInt32(num),
		Price: utils.StrToFloat32(price),
		Name:  name,
		Unit:  unit,
		Pic:   filePath,
		Desc:  desc,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func DelProduct(ctx *gin.Context) {
	id := ctx.PostForm("id")

	service := micro.NewService()
	productService := cyh_product_srv.NewProductsService("cyh_product_srv", service.Client())
	resp, _ := productService.ProductDel(context.TODO(), &cyh_product_srv.ProductDelRequest{Id: utils.StrToInt32(id)})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func ToEditProduct(ctx *gin.Context) {
	id := ctx.Query("id")

	service := micro.NewService()
	productService := cyh_product_srv.NewProductsService("cyh_product_srv", service.Client())
	resp, err := productService.ProductToEdit(context.TODO(), &cyh_product_srv.ProductToEditRequest{Id: utils.StrToInt32(id)})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":       resp.Code,
		"msg":        resp.Msg,
		"product":    resp.Product,
		"img_base64": utils.ImgToBase64(resp.Product.Pic),
	})
}

func EditProduct(ctx *gin.Context) {
	name, price, id := ctx.PostForm("name"), ctx.PostForm("price"), ctx.PostForm("id")
	num, unit := ctx.PostForm("num"), ctx.PostForm("unit")
	desc := ctx.PostForm("desc")

	service := micro.NewService()
	productService := cyh_product_srv.NewProductsService("cyh_product_srv", service.Client())

	file, err := ctx.FormFile("pic")
	if err != nil {
		logger.Error(err)
		resp, _ := productService.ProductDoEdit(context.TODO(), &cyh_product_srv.ProductEditRequest{
			Id:    utils.StrToInt32(id),
			Num:   utils.StrToInt32(num),
			Price: utils.StrToFloat32(price),
			Name:  name,
			Unit:  unit,
			Desc:  desc,
		})

		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}

	filePath := "upload/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	ctx.SaveUploadedFile(file, filePath)

	resp, _ := productService.ProductDoEdit(context.TODO(), &cyh_product_srv.ProductEditRequest{
		Id:    utils.StrToInt32(id),
		Num:   utils.StrToInt32(num),
		Price: utils.StrToFloat32(price),
		Name:  name,
		Unit:  unit,
		Pic:   filePath,
		Desc:  desc,
	})

	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}
