package model

import "gorm.io/gorm"

// 购物车
type Cart struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`

	// 商家 ID
	BossId uint `gorm:"not null"`
	// 购买的数量
	// Quantity int `gorm:"not null"`

	Num uint `gorm:"not null"`
	// 最大可购买数量
	MaxNum uint `gorm:"not null"`
	// 是否下单

	Check bool // IsBuy  bool `gorm:"not null"`
}
