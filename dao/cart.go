package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type CartDao struct {
	*gorm.DB
}

// New CartDao
func NewCartDao(ctx context.Context) *CartDao {
	return &CartDao{
		NewDbCilent(ctx),
	}
}

// New CartDaoByDB
func NewCartDaoByDB(db *gorm.DB) *CartDao {
	return &CartDao{
		db,
	}
}

// 判断是否已经存在
func (dao *CartDao) Exist(userId, productId uint) bool {
	var n int64
	if dao.DB.Model(&model.Cart{}).Where(&model.Cart{UserId: userId, ProductId: productId}).
		Count(&n).Error != nil {
		return false
	}
	return n != 0
}

// GetById 从数据库中获取 Cart
func (dao *CartDao) GeCartById(cid uint) (Cart *model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where(&model.Cart{Model: gorm.Model{ID: cid}}).First(&Cart).Error
	return
}

// add
func (dao *CartDao) Add(cart *model.Cart) (err error) {
	err = dao.DB.Model(&model.Cart{}).Create(cart).Error
	return
}

// Update
func (dao *CartDao) Update(cart *model.Cart) (err error) {
	err = dao.DB.Model(&model.Cart{}).Where(&model.Cart{Model: gorm.Model{ID: cart.ID}}).Updates(cart).Error
	return
}

// Delete
func (dao *CartDao) Delete(uid, productId uint) (err error) {
	m := model.Cart{UserId: uid, ProductId: productId}
	err = dao.DB.Model(&model.Cart{}).Where(&m).Delete(&model.Cart{}).Error
	return
}

// GetAll
func (dao *CartDao) GetAll(uid uint) (list []*model.Cart, err error) {
	err = dao.DB.Model(&model.Cart{}).Where(&model.Cart{UserId: uid}).Find(&list).Error
	return
}
