package common

type BasePage struct {
	Page int `form:"page"`
	Size int `form:"size"` // NOTE 可以写到配置文件中
}
