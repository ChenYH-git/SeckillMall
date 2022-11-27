package user

import (
	"context"
	cyh_user_srv "cyh_project/cyh_user_srv/proto/front_user"
	"cyh_project/cyh_web/utils"
	"net/http"

	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
)

func SendEmail(ctx *gin.Context) {
	email := ctx.PostForm("email")

	ok := utils.VerifyEmailFormat(email)
	if !ok {
		ctx.JSON(http.StatusHTTPVersionNotSupported, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
		return
	}

	service := micro.NewService()
	frontUserService := cyh_user_srv.NewFrontUserService("cyh_user_srv", service.Client())
	resp, _ := frontUserService.FrontUserSendEmail(context.TODO(), &cyh_user_srv.FrontUserMailRequest{Email: email})
	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func FrontUserRegister(ctx *gin.Context) {
	email := ctx.PostForm("email")
	ok := utils.VerifyEmailFormat(email)
	if !ok {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  "邮箱格式不正确",
		})
		return
	}
	captche := ctx.PostForm("captche")
	password := ctx.PostForm("password")
	repassword := ctx.PostForm("repassword")
	if password != repassword {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "两次密码输入不一致",
		})
		return
	}
	service := micro.NewService()
	frontUserService := cyh_user_srv.NewFrontUserService("cyh_user_srv", service.Client())
	resp, err := frontUserService.FrontUserRegister(context.TODO(), &cyh_user_srv.FrontUserRequest{
		Email:      email,
		Code:       captche,
		Password:   password,
		Repassword: repassword,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": resp.Code,
		"msg":  resp.Msg,
	})
}

func FrontUserLogin(ctx *gin.Context) {
	mail := ctx.PostForm("mail")
	password := ctx.PostForm("password")
	if !utils.VerifyEmailFormat(mail) {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  500,
			"token": "邮箱地址错误",
		})
		return
	}

	service := micro.NewService()
	frontUserService := cyh_user_srv.NewFrontUserService("cyh_user_srv", service.Client())
	resp, err := frontUserService.FrontUserLogin(context.TODO(), &cyh_user_srv.FrontUserRequest{
		Email:    mail,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":  resp.Code,
			"token": resp.Msg,
		})
		return
	}

	tokenString, err := utils.GenToken(resp.UserName, utils.FrontUserExpireDuration, utils.FrontUserSecKey)
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
		"username": resp.UserName,
		"token":    tokenString,
	})
}
