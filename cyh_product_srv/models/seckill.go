package models

import "time"

type SecKills struct {
	Id         int
	PId        int
	Num        int
	Status     int
	Price      float32 `gorm:"type:decimal(11,2)"`
	Name       string
	StartTime  time.Time
	EndTime    time.Time
	CreateTime time.Time
}

func (SecKills) TableName() string {
	return "product_seckills"
}
