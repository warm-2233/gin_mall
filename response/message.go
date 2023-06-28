package response

const (
	Success       = "OK"
	Error         = "FAIL"
	InvalidParams = "参数错误"

	ErrorExistUser               = "用户已存在"
	ErrorNotExistUser            = "用户不存在"
	ErrorWrongUsernameOrPassword = "用户名或密码错误"
	ErrorFailEncryption          = "加密失败"
	ErrorAuthToken               = "Token 验证失败"
	ErrorAuthCheckTokenTimeout   = "Token 过期"
	ErrorNotLogin                = "未登录"
	ErrorNotPermission           = "没有权限"

	ErrorUploadFail = "上传失败"

	ErrorSendEmail = "发送邮件失败"

	ErrorFavoriteAlreadyExist = "收藏已存在"
	ErrorFavoriteNotExist     = "收藏不存在"
)
