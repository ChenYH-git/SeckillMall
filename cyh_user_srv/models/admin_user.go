package models

import "time"

type AdminUser struct {
	Id         int
	Status     int
	Username   string
	Password   string
	Desc       string
	CreateTime time.Time
}

func (AdminUser) TableName() string {
	return "admin_user"
}
