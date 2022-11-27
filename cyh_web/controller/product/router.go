package product

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/get_product_list", GetProductList)
	router.POST("/product_add", AddProduct)
	router.POST("/product_del", DelProduct)
	router.POST("/do_product_edit", EditProduct)
	router.GET("/to_product_edit", ToEditProduct)
}
