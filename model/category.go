package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	CategoryName string
	// 父级ID
	// ParentId uint
}
