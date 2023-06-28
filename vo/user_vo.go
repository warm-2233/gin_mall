package vo

import (
	"gin_mall/conf"
	"gin_mall/model"
	"gin_mall/pkg/util"
)

type MoneyVo struct {
	Money float64 `json:"money"`
}
type UserVo struct {
	ID       uint   `json:"id"`
	UserName string `json:"username"`
	Nickname string `json:"nickname"`
	// Type     int    `json:"type"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	Avatar   string `json:"avatar"`
	CreateAt int64  `json:"create_at"`
}

type LoginUserVo struct {
	Token string `json:"token"`
	User  UserVo `json:"user"`
}

// user model to user vo
func UserToVo(user *model.User) *UserVo {
	return &UserVo{
		ID:       user.ID,
		UserName: user.UserName,
		Nickname: user.Nickname,
		Email:    user.Email,
		Status:   user.Status,
		Avatar:   AvatarUrl(user.Avatar),
		CreateAt: user.CreatedAt.Unix(),
	}
}

// 拼接avatar绝对路径
func AvatarUrl(a string) string {
	return conf.PathConf.Host + conf.PathConf.StaticPath[1:] + a
}

// BuildMoney
func BuildMoney(user *model.User, key string) float64 {
	util.Encrypt.SetKey(key)
	return user.Money
}
