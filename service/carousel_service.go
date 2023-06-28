package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/log"
	"gin_mall/response"
	"gin_mall/vo"

	"github.com/gin-gonic/gin"
)

// carouselService 推荐轮播图服务
type carouselService struct {
}

var CarouselService = carouselService{}

func (*carouselService) GetList(ctx context.Context) response.Response {
	resp := response.NewResponse()
	dao := dao.NewCarouselDao(ctx)

	carousels, err := dao.List()
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	resp.Data = gin.H{
		"carousels": vo.ToCarouselVoList(carousels),
		"total":     len(carousels),
	}
	// BUG
	log.InfoLn("not error", err)
	log.Logger.Infoln("12345y6hygtfred")
	log.Logger.Errorln("errrrrrrrrrrrrrrrrrrrrrr")
	return resp
}
