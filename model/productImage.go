package model

import "gorm.io/gorm"

type ProductImage struct {
	gorm.Model
	ProductId uint `gorm:"not null"`
	ImageUrl  string
	// index 产品主图 排序
}
