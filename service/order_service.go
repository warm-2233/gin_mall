package service

import (
	"context"
	"errors"
	"gin_mall/common"
	"gin_mall/dao"
	"gin_mall/dto"
	"gin_mall/model"
	"gin_mall/response"
	"gin_mall/vo"
	"time"

	"gorm.io/gorm"
)

type orderService struct {
}

var OrderService = orderService{}

func (*orderService) Create(ctx context.Context, uid uint, request dto.OrderCreateDto) response.Response {
	resp := response.NewResponse()
	orderDao := dao.NewOrderDao(ctx)
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(request.ProductId)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	addrDao := dao.NewAddressDao(ctx)
	if !addrDao.IsExist(request.AddressId) {
		resp.SetCode(response.ErrorCode)
		resp.Message = "address not exist"
		return resp
	}
	orderID := time.Now().UnixMicro()

	money := product.Price * float64(request.Num)

	order := model.Order{
		UserId:    uid,
		AddressId: request.AddressId,
		ProductId: request.ProductId,
		Num:       request.Num,
		OrderID:   uint(orderID),
		Money:     money,
		Price:     product.Price,
		Type:      1,
	}
	err = orderDao.Create(&order)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	resp.Data = vo.ToOrderVo(&order)
	return resp
}

func (*orderService) List(ctx context.Context, uid uint) response.Response {
	resp := response.NewResponse()
	orderDao := dao.NewOrderDao(ctx)
	orders, err := orderDao.List(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	ordersVo := []*vo.OrderVo{}
	for _, v := range orders {
		ordersVo = append(ordersVo, vo.ToOrderVo(v))
	}

	resp.Data = ordersVo

	return resp
}

func (*orderService) Pay(ctx context.Context, uid uint, orderID uint) response.Response {
	resp := response.NewResponse()
	orderDao := dao.NewOrderDao(ctx)
	order, err := orderDao.GetByOrderId(orderID)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if order.UserId != uid {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("Order not found")
		return resp
	}
	if order.Type != common.OrderTypePending {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("Order not found!")
		return resp
	}
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(order.ProductId)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	userDao := dao.NewUserDao(ctx)
	boss, err := userDao.GetById(product.BossId)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	user, err := userDao.GetById(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	err = orderDao.DB.Transaction(func(tx *gorm.DB) error {
		if user.Money-order.Money < 0.0 {
			return errors.New("金额不足")
		}
		user.Money = user.Money - order.Money
		err := tx.Model(&model.User{}).Where("id=?", user.ID).Updates(&user).Error
		if err != nil {
			return err
		}
		boss.Money = boss.Money + order.Money
		err = tx.Model(&model.User{}).Where("id=?", boss.ID).Updates(&boss).Error
		if err != nil {
			return err
		}
		if product.Num-order.Num < 0 {
			return errors.New("库存不足")
		}
		product.Num = product.Num - order.Num
		err = tx.Model(&model.Product{}).Where("id=?", product.ID).Save(&product).Error
		if err != nil {
			return err
		}
		order.Type = common.OrderTypePaid
		err = tx.Model(&model.Order{}).Where("id=?", order.ID).Updates(&order).Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	return resp
}
