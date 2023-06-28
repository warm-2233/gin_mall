package model

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId    uint `gorm:"not null"`
	ProductId uint `gorm:"not null"`
	AddressId uint `gorm:"not null"`
	Num       int
	OrderID   uint
	Price     float64

	// 1 未支付， 2 已支付， 3 已完成
	Type int
	// 总金额
	Money float64
}

/**
1 待付款
2 代发货，已支付
3 已发货
4 待收货
5 已完成
6 已取消
*/

// type Order struct {
//     OrderID              int
//     OrderTime            time.Time
//     PaymentTime          time.Time
//     ProductID            int
//     ProductName          string
//     Quantity             int
//     Price                float64
//     ShippingAddress      string
//     RecipientName        string
//     RecipientPhoneNumber string
//     OrderStatus          string
// }

// 订单号（Order ID）
// 下单时间（Order time）
// 付款时间（Payment time）
// 商品ID（Product ID）
// 商品名称（Product Name）
// 商品数量（Quantity）
// 商品价格（Price）
// 收货地址（Shipping address）
// 收货人姓名（Recipient name）
// 收货人手机号码（Recipient phone number）
// 订单状态（Order status）
// CREATE TABLE orders (
//     order_id INT PRIMARY KEY,
//     order_time DATETIME,
//     payment_time DATETIME,
//     product_id INT,
//     product_name VARCHAR(255),
//     quantity INT,
//     price DECIMAL(10,2),
//     shipping_address VARCHAR(255),
//     recipient_name VARCHAR(255),
//     recipient_phone_number VARCHAR(20),
//     order_status VARCHAR(20)
// );

// 订单状态
