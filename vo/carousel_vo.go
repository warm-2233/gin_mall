package vo

import (
	"gin_mall/model"
	"gin_mall/pkg/util"
)

type CarouselVo struct {
	ImageUrl   string `json:"image_url,omitempty"`
	ProductUrl string `json:"product_url,omitempty"`
}

type CarouselVoList []CarouselVo

func ToCarouselVoList(carousels []model.Carousel) CarouselVoList {
	var carouselVos CarouselVoList
	for _, carousel := range carousels {
		carouselVos = append(carouselVos, CarouselVo{
			ImageUrl:   carousel.ImageUrl,
			ProductUrl: util.GetProductUrl(carousel.ProductId),
		})
	}
	return carouselVos
}
