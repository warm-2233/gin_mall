package model

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Uid     uint   `gorm:"not null"`
	Name    string `gorm:"type:varchar(20) not null"`
	Phone   string `gorm:"type:varchar(11) not null"`
	Address string `gorm:"type:varchar(256) not null"`
}
