package model

import "gorm.io/gorm"

// 轮播图
type Carousel struct {
	gorm.Model
	ImageUrl string
	// 产品ID
	ProductId uint `gorm:"not null"`
}
