package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type ProductImageDao struct {
	*gorm.DB
}

func NewProductImageDao(ctx context.Context) *ProductImageDao {
	return &ProductImageDao{
		DB: NewDbCilent(ctx),
	}
}

func NewProductImageDaoByDB(db *gorm.DB) *ProductImageDao {
	return &ProductImageDao{
		db,
	}
}

func (dao *ProductImageDao) Create(img *model.ProductImage) error {
	return dao.DB.Model(&model.ProductImage{}).Create(img).Error
}
