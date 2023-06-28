package dto

import "gin_mall/common"

// TODO 参数必要 验证
type ProductDto struct {
	Name     string  `json:"name" form:"name"`
	Category uint    `json:"category" form:"category"`
	Title    string  `json:"title" form:"title"`
	Info     string  `json:"info" form:"info"`
	ImgURL   string  `json:"img_url" form:"img_url"`
	Price    float64 `json:"price" form:"price"`
	Discount string  `json:"discount" form:"discount"`
	OnSale   bool    `json:"on_sale" form:"on_sale"`
	Stock    int     `json:"stock" form:"stock"`
}

type ProductListDto struct {
	Info     string `json:"info" form:"info"`
	Category uint   `json:"category" form:"category"`
	common.BasePage
}
