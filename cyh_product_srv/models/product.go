package models

import "time"

type Products struct {
	Id         int
	Num        int
	Price      float32 `gorm:"type: decimal(11, 2)"`
	Name       string
	Unit       string
	Pic        string
	Desc       string
	CreateTime time.Time
	SecKills   []SecKills `gorm:"ForeignKey: PId;AssociationForeignKey: Id"`
}

func (Products) TableName() string {
	return "products"
}
