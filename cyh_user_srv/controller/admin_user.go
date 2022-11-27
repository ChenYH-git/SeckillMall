package controller

import (
	"context"
	"cyh_project/cyh_user_srv/data_source"
	"cyh_project/cyh_user_srv/models"
	cyh_user_srv "cyh_project/cyh_user_srv/proto/admin_user"
	"cyh_project/cyh_user_srv/utils"
	"errors"
	"strconv"
)

type AdminUser struct{}

func (a *AdminUser) AdminUserLogin(ctx context.Context, in *cyh_user_srv.AdminUserRequest, out *cyh_user_srv.AdminUserResponse) error {
	username, password := in.Username, in.Password
	adminUser, cnt := &models.AdminUser{}, 0
	saltPwd := utils.Md5Pwd(password)

	data_source.Db.Where("username = ?", username).Where("password = ?", saltPwd).Find(adminUser).Count(&cnt)
	if cnt < 1 {
		out.Msg = "管理员用户名或密码错误"
		out.Code = 500
		return errors.New("用户名或密码错误")
	}

	out.Msg = "登录成功"
	out.Code = 200
	out.UserName = username
	return nil
}

func (a *AdminUser) FrontUserList(ctx context.Context, in *cyh_user_srv.FrontUsersRequest, out *cyh_user_srv.FrontUsersResponse) error {
	currentPage, pageSize := in.CurrentPage, in.PageSize

	var users []models.FrontUser
	var user models.FrontUser
	res := data_source.Db.Limit(pageSize).Offset((currentPage - 1) * pageSize).Find(&users)

	if res.Error != nil {
		out.Code = 500
		out.Msg = "查询用户列表失败"
		return errors.New("用户列表查询失败")
	}

	total := int32(data_source.Db.Find(&user).RowsAffected)
	var frontUsers []*cyh_user_srv.User
	for _, user := range users {
		frontUser := &cyh_user_srv.User{
			Email:      user.Email,
			Desc:       user.Desc,
			Status:     strconv.FormatInt(int64(user.Status), 10),
			CreateTime: user.CreateTime.Format("2006-01-02 15:04:05"),
		}
		frontUsers = append(frontUsers, frontUser)
	}

	out.Code = 200
	out.Msg = "获取用户列表成功"
	out.FrontUsers = frontUsers
	out.Total = total
	out.PageSize = pageSize
	out.CurrentPage = currentPage
	return nil
}
