package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type AddressDao struct {
	*gorm.DB
}

func NewAddressDao(ctx context.Context) *AddressDao {
	return &AddressDao{
		DB: NewDbCilent(ctx),
	}
}

func NewAddressDaoByDB(db *gorm.DB) *AddressDao {
	return &AddressDao{
		db,
	}
}

// 判断是否存在
func (dao *AddressDao) IsExist(aid uint) bool {
	var n int64
	if dao.DB.Model(&model.Address{}).Where(&model.Address{Model: gorm.Model{ID: aid}}).
		Count(&n).Error != nil {
		return false
	}
	return n != 0
}

func (dao *AddressDao) Create(a *model.Address) error {
	return dao.DB.Model(&model.Address{}).Create(a).Error
}

func (dao *AddressDao) Update(a *model.Address) error {
	return dao.DB.Model(&model.Address{}).Where(&model.Address{Model: gorm.Model{ID: a.ID}}).Updates(a).Error
}

func (dao *AddressDao) Delete(a *model.Address) error {
	return dao.DB.Model(&model.Address{}).Delete(a).Error
}
func (dao *AddressDao) Get(aid uint) (addr *model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where(&model.Address{Model: gorm.Model{ID: aid}}).First(&addr).Error
	return
}

func (dao *AddressDao) List(uid uint) (addrs []*model.Address, err error) {
	err = dao.DB.Model(&model.Address{}).Where(&model.Address{Uid: uid}).Find(&addrs).Error
	return
}
