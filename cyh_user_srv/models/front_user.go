package models

import "time"

type FrontUser struct {
	Id         int
	Status     int
	Email      string `gorm:"unique;not null"`
	Password   string
	Desc       string
	CreateTime time.Time
}

func (FrontUser) TableName() string {
	return "front_user"
}
