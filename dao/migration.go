package dao

import (
	"context"
	"gin_mall/common"
	"gin_mall/model"

	"gorm.io/gorm"
)

func Migration() {
	_db := common.GetDb()
	err := _db.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(
			&model.Address{},
			&model.Admin{},
			&model.Carousel{},
			&model.Cart{},
			&model.Category{},
			&model.Favorite{},
			&model.Notice{},
			&model.Order{},
			&model.ProductImage{},
			&model.Product{},
			&model.User{},
		)

	if err != nil {
		panic(err)
	}
}

func NewDbCilent(ctx context.Context) *gorm.DB {
	return common.NewDbCilent(ctx)
}
