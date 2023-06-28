package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type FavoriteDao struct {
	*gorm.DB
}

func NewFavoriteDao(ctx context.Context) *FavoriteDao {
	return &FavoriteDao{
		NewDbCilent(ctx),
	}
}

// 是否已经收藏
func (dao *FavoriteDao) IsFavorite(uid, pid uint) (bool, error) {
	var n int64
	err := dao.DB.Model(&model.Favorite{}).Where(&model.Favorite{UserId: uid, ProductId: pid}).Count(&n).Error
	if err != nil {
		return false, err
	}
	return n > 0, nil

}

func (dao *FavoriteDao) List(uid uint, page, pageSize int) (categories []*model.Favorite, err error) {
	// Preload 预处理，外键关联自动查询到struct中
	err = dao.DB.Model(&model.Favorite{}).Preload("User").Preload("Product").
		Where(&model.Favorite{UserId: uid}).Offset(page * pageSize).Limit(pageSize).
		Find(&categories).Error
	return
}

func (dao *FavoriteDao) Create(uid, pid uint) error {
	return dao.DB.Model(&model.Favorite{}).Create(&model.Favorite{UserId: uid, ProductId: pid, BossId: uid}).Error
}

func (dao *FavoriteDao) Delete(uid, pid uint) error {
	return dao.DB.Model(&model.Favorite{}).Delete(&model.Favorite{}, &model.Favorite{UserId: uid, ProductId: pid}).Error
}
