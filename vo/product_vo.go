package vo

import (
	"context"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
)

type ProductVo struct {
	Id         uint    `json:"id"`
	Name       string  `json:"name"`
	Category   uint    `json:"category"`
	Title      string  `json:"title"`
	Info       string  `json:"info"`
	ImageURL   string  `json:"image_url"`
	Price      float64 `json:"price"`
	Discount   string  `json:"discount"`
	Stock      int     `json:"stock"`
	OnSale     bool    `json:"on_sale"`
	CreatedAt  int64   `json:"created_at"`
	BossId     uint    `json:"boss_id"`
	BossAvatar string  `json:"boss_avatar"`
	BossName   string  `json:"boss_name"`
	View       int64   `json:"view"`

	// 商品详情
	Images []string `json:"images,omitempty"`
}

func ToProductVo(product *model.Product) *ProductVo {
	return &ProductVo{
		Id:         product.ID,
		Name:       product.Name,
		Category:   product.Category,
		Title:      product.Title,
		Info:       product.Info,
		ImageURL:   ProductUrl(product.ImageUrl),
		Price:      product.Price,
		Discount:   product.DiscountPrice,
		Stock:      product.Num,
		OnSale:     product.OnSale,
		CreatedAt:  product.CreatedAt.Unix(),
		BossId:     product.BossId,
		BossAvatar: product.BossAvatar,
		BossName:   product.BossName,
		View:       product.View(),
		Images:     GetImages(product),
	}
}

type Products struct {
	Category uint    `json:"category"`
	Title    string  `json:"title"`
	Info     string  `json:"info"`
	ImageURL string  `json:"image_url"`
	Price    float64 `json:"price"`
	Discount string  `json:"discount"`
	Stock    int     `json:"stock"`
}
type ProductListVo struct {
	Total    int        `json:"total"`
	Products []Products `json:"products"`
}

func ToProductListVo(products []*model.Product) ProductListVo {
	var productVos ProductListVo
	for _, product := range products {
		productVos.Products = append(productVos.Products, Products{
			Category: product.Category,
			Title:    product.Title,
			Info:     product.Info,
			ImageURL: product.ImageUrl,
			Price:    product.Price,
			Discount: product.DiscountPrice,
			Stock:    product.Num,
		})
	}
	return productVos
}

func ProductUrl(a string) string {
	return conf.PathConf.Host + conf.PathConf.StaticPath + a
}

func GetImages(product *model.Product) []string {
	dao := dao.NewProductDao(context.Background())
	images := dao.ImageList(product)

	for i := 0; i < len(images); i++ {
		images[i] = ProductUrl(images[i])
	}
	return images

}
