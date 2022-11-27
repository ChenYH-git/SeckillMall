package controller

import (
	"context"
	"cyh_project/cyh_user_srv/data_source"
	"cyh_project/cyh_user_srv/models"
	cyh_user_srv "cyh_project/cyh_user_srv/proto/front_user"
	"cyh_project/cyh_user_srv/utils"
	"time"

	"errors"

	"github.com/patrickmn/go-cache"
)

type FrontUser struct{}

var c = cache.New(60*time.Second, 10*time.Second)

func (f *FrontUser) FrontUserRegister(ctx context.Context, in *cyh_user_srv.FrontUserRequest, out *cyh_user_srv.FrontUserResponse) error {
	email := in.Email
	captche := in.Code
	password := in.Password

	code, ok := c.Get(email)
	if !ok {
		out.Code = 500
		out.Msg = "注册失败，请重新尝试"
		return errors.New("验证码获取失败")
	}
	if code != captche {
		out.Code = 500
		out.Msg = "邮箱验证码不正确"
		return errors.New("验证码不一致")
	}
	pwdSalt := utils.Md5Pwd(password)
	frontUser := models.FrontUser{
		Email:      email,
		Password:   pwdSalt,
		Status:     1,
		CreateTime: time.Now(),
	}
	data_source.Db.Create(&frontUser)
	out.Code = 200
	out.Msg = "注册成功"
	return nil
}

func (f *FrontUser) FrontUserSendEmail(ctx context.Context, in *cyh_user_srv.FrontUserMailRequest, out *cyh_user_srv.FrontUserResponse) error {
	email := in.Email

	frontUser := &models.FrontUser{}
	if data_source.Db.Where("email = ?", email).Find(frontUser).RowsAffected >= 1 {
		out.Code = 500
		out.Msg = "邮箱已注册，请使用未注册的邮箱"
		return nil
	}

	randNum := utils.GenEmailCode(6)
	utils.SendEmail(email, randNum)

	c.Set(email, randNum, cache.DefaultExpiration)
	out.Code = 200
	out.Msg = "邮件发送成功"
	return nil
}

func (f *FrontUser) FrontUserLogin(ctx context.Context, in *cyh_user_srv.FrontUserRequest, out *cyh_user_srv.FrontUserResponse) error {
	email, password := in.Email, in.Password
	saltPwd := utils.Md5Pwd(password)
	frontUser := models.FrontUser{}
	if data_source.Db.Where("email = ? AND password = ?", email, saltPwd).Find(&frontUser).RowsAffected < 1 {
		out.Code = 500
		out.Msg = "用户名或密码错误"
		return errors.New("登录校验失败")
	}
	out.Code = 200
	out.Msg = "登录成功"
	out.UserName = email
	return nil
}
