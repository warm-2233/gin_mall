package v1

import (
	"gin_mall/response"
	"gin_mall/service"

	"github.com/gin-gonic/gin"
)

//	@Tags		上传文件
//	@Produce	json
//	@Summary	上传文件
//	@Param		Authorization	header		string	true	"登录凭证"
//	@Param		file			formData	file	true	"文件"
//	@Success	200				{object}	response.Response{data=service.UploadVo}
//	@Router		/api/v1/upload  [post]
func Upload(c *gin.Context) {
	uid := c.GetUint("ID")
	ftype := c.Query("type")
	if (ftype == "") || (ftype != "product") {
		ftype = "image"
	}
	_, fileHeader, err := c.Request.FormFile("file")

	if err != nil {
		response.FailResponse(c, response.InvalidParamsCode, err.Error())
		return
	}

	resp := service.UploadService.UploadToLocalStatic(uid, fileHeader, ftype)
	response.SuccessResponse(c, resp)
}
