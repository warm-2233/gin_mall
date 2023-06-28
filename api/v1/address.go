package v1

import (
	"gin_mall/dto"
	"gin_mall/response"
	"gin_mall/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListAddresses(c *gin.Context) {
	uid := c.GetUint("ID")
	resp := service.Address.ListAddresses(c.Request.Context(), uid)
	response.SuccessResponse(c, resp)

}
func GetAddresses(c *gin.Context) {
	aid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.InvalidParamsResponse(c, "缺少必要参数")
		return
	}

	uid := c.GetUint("ID")
	r := service.Address.GetAddresses(c.Request.Context(), uint(aid), uid)
	response.SuccessResponse(c, r)

}
func AddressCreate(c *gin.Context) {
	uid := c.GetUint("ID")
	request := dto.AddressCreateDto{}
	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, "缺少必要参数")
		return
	}
	resp := service.Address.AddressCreate(c.Request.Context(), uid, request)
	response.SuccessResponse(c, resp)

}
func AddressUpdate(c *gin.Context) {
	request := dto.AddressUpdateDto{}
	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, "缺少必要参数")
		return
	}

	uid := c.GetUint("ID")

	r := service.Address.AddressUpdate(c.Request.Context(), request, uid)
	response.SuccessResponse(c, r)
}
func AddressDelete(c *gin.Context) {
	aid, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.InvalidParamsResponse(c, "缺少必要参数")
		return
	}

	uid := c.GetUint("ID")

	r := service.Address.AddressDelete(c.Request.Context(), uint(aid), uid)
	response.SuccessResponse(c, r)
}
