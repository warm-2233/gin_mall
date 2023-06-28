package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type CarouselDao struct {
	*gorm.DB
}

// New CarouselDao
func NewCarouselDao(ctx context.Context) *CarouselDao {
	return &CarouselDao{
		NewDbCilent(ctx),
	}
}

// New CarouselDaoByDB
func NewCarouselDaoByDB(db *gorm.DB) *CarouselDao {
	return &CarouselDao{
		db,
	}
}

// GetById 从数据库中获取 Carousel
func (dao *CarouselDao) List() (carousel []model.Carousel, err error) {
	err = dao.DB.Model(model.Carousel{}).Find(&carousel).Error
	return
}
