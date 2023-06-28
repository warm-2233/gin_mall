package dto

type UserRegisterDto struct {
	Nickname string `json:"nickname"`
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserLoginDto struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateDto struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
