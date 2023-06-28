package v1

import (
	"gin_mall/response"
	"gin_mall/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendEmail(c *gin.Context) {
	var request service.SendEmailService
	uid := c.GetUint("ID")

	if err := c.ShouldBindJSON(&request); err != nil {
		if err != nil {
			response.InvalidParamsResponse(c, err.Error())
			return
		}
	}
	resp := request.SendEmail(c.Request.Context(), uid)
	c.JSON(http.StatusOK, resp)
	return
}

func ValidateEmail(c *gin.Context) {
	var request service.ValidateEmailService
	if err := c.ShouldBindJSON(&request); err != nil {
		if err != nil {
			response.InvalidParamsResponse(c, err.Error())
			return
		}
	}

	resp := request.ValidateEmail(c.Request.Context(), c.GetHeader("Authorization"))
	c.JSON(http.StatusOK, resp)
	return
}
