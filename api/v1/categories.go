package v1

import (
	"gin_mall/response"
	"gin_mall/service"

	"github.com/gin-gonic/gin"
)

// ListCategories
func ListCategories(c *gin.Context) {
	resp := service.CategoryService.List(c.Request.Context())
	response.SuccessResponse(c, resp)
}
