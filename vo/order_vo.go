package vo

import (
	"gin_mall/common"
	"gin_mall/model"
)

type OrderVo struct {
	UserId    uint    `json:"user_id"`
	Product   uint    `json:"product"`
	AddressId uint    `json:"address_id"`
	Num       int     `json:"num"`
	OrderID   uint    `json:"order_id"`
	Price     float64 `json:"price"`

	// 1 未支付， 2 已支付， 3 已完成
	Type string `json:"type"`
	// 总金额
	Money float64 `json:"money"`
}

func ToOrderVo(order *model.Order) *OrderVo {
	return &OrderVo{
		UserId:    order.UserId,
		Product:   order.ProductId,
		AddressId: order.AddressId,
		Num:       order.Num,
		OrderID:   order.OrderID,
		Price:     order.Price,
		Type:      common.OrderTypeToString(order.Type),
		Money:     order.Money,
	}
}
