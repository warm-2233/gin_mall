package v1

import (
	"gin_mall/dto"
	"gin_mall/response"
	"gin_mall/service"

	"github.com/gin-gonic/gin"
)

//	@Tags		用户相关路由
//	@Produce	json
//	@Summary	用户注册
//	@Param		UserRegisterDto	body		dto.UserRegisterDto	true	"注册参数"
//	@Success	200				{object}	response.Response
//	@Router		/api/v1/user/register  [post]
func UserRegister(c *gin.Context) {
	var request dto.UserRegisterDto

	err := c.ShouldBind(&request)
	if err != nil {
		response.FailResponse(c, response.InvalidParamsCode, err.Error())
		return
	}

	if request.Nickname == "" {
		request.Nickname = request.UserName

	}
	resp := service.UserService.Register(c.Request.Context(), request)
	response.SuccessResponse(c, resp)
}

//	@Tags		用户相关路由
//	@Produce	json
//	@Summary	用户登录
//	@Param		UserLoginDto	body		dto.UserLoginDto	true	"登录参数"
//	@Success	200				{object}	response.Response{data=vo.LoginUserVo}
//	@Router		/api/v1/user/login  [post]
func UserLogin(c *gin.Context) {
	var request dto.UserLoginDto

	err := c.ShouldBind(&request)
	if err == nil {
		resp := service.UserService.Login(c.Request.Context(), request)
		response.SuccessResponse(c, resp)
		return
	}
	response.FailResponse(c, response.InvalidParamsCode, err.Error())
}

//	@Tags		用户相关路由
//	@Produce	json
//	@Summary	更新用户信息
//	@Param		Authorization	header		string				true	"登录凭证"
//	@Param		UserUpdateDto	body		dto.UserUpdateDto	false	"更新用户信息参数"
//	@Success	200				{object}	response.Response{data=vo.UserVo}
//	@Router		/api/v1/user  [put]
func UserUpdate(c *gin.Context) {
	var request dto.UserUpdateDto

	uid := c.GetUint("ID")

	err := c.ShouldBind(&request)
	if err == nil {
		resp := service.UserService.Update(c.Request.Context(), uid, request)
		response.SuccessResponse(c, resp)
		return
	}
	response.FailResponse(c, response.InvalidParamsCode, err.Error())
}

//	@Tags		用户相关路由
//	@Produce	json
//	@Summary	上传头像
//	@Param		Authorization	header		string	true	"登录凭证"
//	@Param		file			formData	file	true	"头像文件"
//	@Success	200				{object}	response.Response{data=vo.UserVo}
//	@Router		/api/v1/avatar  [post]
func UserUploadAvatar(c *gin.Context) {
	uid := c.GetUint("ID")

	_, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.FailResponse(c, response.InvalidParamsCode, err.Error())
		return
	}

	if err == nil {
		resp := service.UserService.UploadAvatar(c.Request.Context(), uid, fileHeader)
		response.SuccessResponse(c, resp)
		return
	}
	response.FailResponse(c, response.InvalidParamsCode, err.Error())
}

//	@Tags		用户相关路由
//	@Produce	json
//	@Summary	显示金额
//	@Param		Authorization	header		string	true	"登录凭证"
//	@Success	200				{object}	response.Response{data=vo.MoneyVo}
//	@Router		/api/v1/money [post]
func ShowMoney(c *gin.Context) {
	uid := c.GetUint("ID")
	resp := service.UserService.ShowMoney(c.Request.Context(), uid)
	response.SuccessResponse(c, resp)
}
