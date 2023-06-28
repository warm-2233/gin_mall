package model

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	UserName string
	Password string
	Avatar   string
}
