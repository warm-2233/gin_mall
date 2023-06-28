package v1

import (
	"gin_mall/common"
	"gin_mall/response"
	"gin_mall/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ListFavorites
func ListFavorites(c *gin.Context) {
	uid := c.GetUint("ID")
	if uid == 0 {
		response.InvalidParamsResponse(c, "")
		return
	}
	request := common.BasePage{}
	err := c.ShouldBind(&request)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}
	if request.Size == 0 {
		request.Size = 10
	}
	resp := service.FavoriteService.List(c.Request.Context(), uid, request)
	c.JSON(http.StatusOK, resp)
}

func FavoriteCreate(c *gin.Context) {
	uid := c.GetUint("ID")
	if uid == 0 {
		response.InvalidParamsResponse(c, "")
		return
	}
	pid := c.Query("pid")
	if pid == "" {
		response.InvalidParamsResponse(c, "")
		return
	}
	intPid, err := strconv.Atoi(pid)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}
	resp := service.FavoriteService.AddFavorite(c.Request.Context(), uid, uint(intPid))
	c.JSON(http.StatusOK, resp)
}

// FavoriteDelete
func FavoriteDelete(c *gin.Context) {
	uid := c.GetUint("ID")
	if uid == 0 {
		response.InvalidParamsResponse(c, "")
		return
	}
	pid := c.Query("pid")
	if pid == "" {
		response.InvalidParamsResponse(c, "")
		return
	}
	intPid, err := strconv.Atoi(pid)
	if err != nil {
		response.InvalidParamsResponse(c, err.Error())
		return
	}
	resp := service.FavoriteService.DeleteFavorite(c.Request.Context(), uid, uint(intPid))
	c.JSON(http.StatusOK, resp)
}
