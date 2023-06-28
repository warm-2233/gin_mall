package model

import "gorm.io/gorm"

// NOTE
// 一般来说，电商中一个商品的收藏夹的表结构可能包括以下字段：
// 用户ID：收藏该商品的用户ID。
// 商品ID：被收藏的商品ID。
// 收藏时间：用户收藏该商品的时间。
// 删除状态：记录用户是否已将该商品从收藏夹中删除。
// 其他商品信息：包括商品名称、价格、图片等信息，这些信息可以通过联表查询获取。
// 具体的表结构可能因具体业务需求而有所不同，以上仅为一般性参考。

type Favorite struct {
	gorm.Model

	// User 外键
	User      *User    `gorm:"ForeignKey:UserId"`
	UserId    uint     `gorm:"not null"`
	Product   *Product `gorm:"ForeignKey:ProductId"`
	ProductId uint     `gorm:"not null"`
	Boss      *User    `gorm:"ForeignKey:BossId"`
	BossId    uint     `gorm:"not null"`
}
