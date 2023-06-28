package v1

import (
	"gin_mall/response"
	"gin_mall/service"

	"github.com/gin-gonic/gin"
)

// ListCarousels
func ListCarousels(c *gin.Context) {
	resp := service.CarouselService.GetList(c.Request.Context())
	response.SuccessResponse(c, resp)
}
