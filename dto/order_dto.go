package dto

type OrderCreateDto struct {
	ProductId uint `json:"product_id" binding:"required"`
	AddressId uint `json:"address_id" binding:"required"`
	Num       int  `json:"num" binding:"required"`
}

type OrderUpdateDto struct {
	AddressId uint `json:"address_id"`
	Num       int  `json:"num"`
}

type OrderPayDto struct {
	OrderId uint `json:"order_id" binding:"required"`
}
