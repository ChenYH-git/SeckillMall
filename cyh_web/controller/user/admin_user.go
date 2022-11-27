package user

import (
	"context"
	cyh_user_srv "cyh_project/cyh_user_srv/proto/admin_user"
	"cyh_project/cyh_web/utils"
	"net/http"

	"go-micro.dev/v4"

	"github.com/gin-gonic/gin"
)

func AdminLogin(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	service := micro.NewService()
	adminUserService := cyh_user_srv.NewAdminUserService("cyh_user_srv", service.Client())
	resp, err := adminUserService.AdminUserLogin(context.TODO(), &cyh_user_srv.AdminUserRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code": resp.Code,
			"msg":  resp.Msg,
		})
		return
	}
	adminToken, err := utils.GenToken(username, utils.AdminUserExpireDuration, utils.AdminUserSecKey)
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
		"user_name":   resp.UserName,
		"admin_token": adminToken,
	})
}

func FrontUserList(ctx *gin.Context) {
	page, pageSize := ctx.DefaultQuery("currentPage", "1"), ctx.DefaultQuery("pageSize", "5")

	service := micro.NewService()
	adminUserService := cyh_user_srv.NewAdminUserService("cyh_user_srv", service.Client())
	resp, err := adminUserService.FrontUserList(context.TODO(), &cyh_user_srv.FrontUsersRequest{
		CurrentPage: utils.StrToInt32(page),
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
		"front_users":  resp.FrontUsers,
		"total":        resp.Total,
		"current_page": resp.CurrentPage,
		"page_size":    resp.PageSize,
	})
}
