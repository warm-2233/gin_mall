package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type OrderDao struct {
	*gorm.DB
}

// New OrderDao
func NewOrderDao(ctx context.Context) *OrderDao {
	return &OrderDao{
		NewDbCilent(ctx),
	}
}

// New OrderDaoByDB
func NewOrderDaoByDB(db *gorm.DB) *OrderDao {
	return &OrderDao{
		db,
	}
}

// get by order id
func (od *OrderDao) GetByOrderId(oid uint) (o *model.Order, err error) {
	err = od.DB.Model(&model.Order{}).Where(&model.Order{OrderID: oid}).First(&o).Error
	return
}

func (dao *OrderDao) Create(order *model.Order) error {
	return dao.DB.Create(order).Error
}

func (dao *OrderDao) List(uid uint) (orders []*model.Order, err error) {
	err = dao.DB.Model(&model.Order{}).Where(&model.Order{UserId: uid}).Find(&orders).Error
	return
}
