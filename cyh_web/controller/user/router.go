package user

import (
	"cyh_project/cyh_web/middleware"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.POST("/send_email", SendEmail)
	router.POST("/front_user_register", FrontUserRegister)
	router.POST("/front_user_login", FrontUserLogin)
	router.POST("/admin_login", AdminLogin)
	router.GET("/get_front_users", middleware.JwtTokenCheckOut, FrontUserList)
}
