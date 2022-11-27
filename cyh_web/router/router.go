package all_router

import (
	"cyh_project/cyh_web/controller/product"
	"cyh_project/cyh_web/controller/seckill"
	"cyh_project/cyh_web/controller/user"
	"cyh_project/cyh_web/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(router *gin.Engine) {
	userGroup := router.Group("/user")
	productGroup := router.Group("/product")
	seckillGroup := router.Group("/seckill")

	user.Router(userGroup)
	productGroup.Use(middleware.JwtTokenCheckOut)
	product.Router(productGroup)
	seckill.Router(seckillGroup)
}
