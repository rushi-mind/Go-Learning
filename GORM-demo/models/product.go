package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Category string  `gorm:"column:category" json:"category"`
	Name     string  `gorm:"column:name;not null" json:"name"`
	Price    float32 `gorm:"column:price;not null" json:"price"`
}
