package seckill

import (
	"cyh_project/cyh_web/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.GET("/get_seckill_list", middleware.JwtTokenCheckOut, GetSecKillList)
	router.GET("/get_products", middleware.JwtTokenCheckOut, GetProducts)
	router.POST("/seckill_add", middleware.JwtTokenCheckOut, AddSecKill)
	router.POST("/seckill_del", middleware.JwtTokenCheckOut, DelSecKill)
	router.POST("/seckill_do_edit", middleware.JwtTokenCheckOut, EditSecKill)
	router.GET("/seckill_to_edit", middleware.JwtTokenCheckOut, ToEditSecKill)

	// 前端列表
	router.GET("/front/get_seckill_list", GetFrontSeckillList)
	// 前端详情
	router.GET("/front/seckill_detail", middleware.JwtFrontTokenCheckOut, SecKillDetail)

	// 秒杀接口
	router.POST("/front/seckill", middleware.JwtFrontTokenCheckOut, SecKill)

	// 获取下单结果
	router.GET("/front/get_seckill_result", middleware.JwtFrontTokenCheckOut, GetSeckillResult)
}
