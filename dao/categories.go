package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type CategoryDao struct {
	*gorm.DB
}

func NewCategoryDao(ctx context.Context) *CategoryDao {
	return &CategoryDao{
		NewDbCilent(ctx),
	}
}

func (dao *CategoryDao) List() (categories []*model.Category, err error) {
	err = dao.DB.Model(&model.Category{}).Find(&categories).Error
	return
}
