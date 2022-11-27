package middleware

import (
	"cyh_project/cyh_web/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtTokenCheckOut(ctx *gin.Context) {
	authHeader := ctx.Request.Header.Get("Authorization")
	if authHeader == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		ctx.Abort()
		return
	}

	auths := strings.SplitN(authHeader, " ", 2)
	bearer, token := auths[0], auths[1]
	if len(token) == 0 || bearer != "Bearer" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "token格式错误",
		})
		ctx.Abort()
		return
	}

	user, err := utils.AuthToken(token, utils.AdminUserSecKey)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "invalid token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("adminUserName", user.Username)
	ctx.Next()
}

func JwtFrontTokenCheckOut(ctx *gin.Context) {

	authHeader := ctx.Request.Header.Get("Authorization")

	if authHeader == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带token",
		})
		ctx.Abort()
		return
	}

	auths := strings.Split(authHeader, " ")

	bearer := auths[0]
	token := auths[1]

	if len(token) == 0 || len(bearer) == 0 {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "请携带正确格式的token",
		})
		ctx.Abort()
		return
	}

	user, err := utils.AuthToken(token, utils.FrontUserSecKey)

	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 401,
			"msg":  "无效的token",
		})
		ctx.Abort()
		return
	}

	ctx.Set("frontUserEmail", user.Username)
	ctx.Next()
}
