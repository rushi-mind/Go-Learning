package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model
	Name       string `gorm:"column:name" json:"name"`
	ShopNumber int    `gorm:"column:shopNumber" json:"shopNumber"`
}
