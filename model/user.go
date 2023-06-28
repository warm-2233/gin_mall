package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string
	Nickname       string
	Status         string
	Avatar         string
	Money          float64
	Authority      int
}

const PasswordCost = 12 // 密码加密长度

func (u *User) SetPassword(s string) error {
	b, err := bcrypt.GenerateFromPassword([]byte(s), PasswordCost)
	if err != nil {
		return err
	}

	u.PasswordDigest = string(b)

	return nil
}

// 判断密码是否正确
func (u *User) CheckPassword(s string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordDigest), []byte(s))
	return err == nil
}
