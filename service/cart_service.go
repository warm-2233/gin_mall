package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/response"
)

type cartService struct{}

var CartService = cartService{}

func (*cartService) AddCart(ctx context.Context, uid uint, productId uint) response.Response {
	resp := response.NewResponse()
	cartDao := dao.NewCartDao(ctx)
	cart := model.Cart{
		UserId:    uid,
		ProductId: productId,
	}

	if cartDao.Exist(uid, productId) {
		resp.SetCode(response.ErrorCode)
		resp.Message = "cart already exist"
		return resp
	}

	err := cartDao.Add(&cart)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	return resp
}

func (*cartService) GetCart(ctx context.Context, uid uint) response.Response {
	resp := response.NewResponse()
	cartDao := dao.NewCartDao(ctx)
	list, err := cartDao.GetAll(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = list

	return resp
}

func (*cartService) DeleteCart(ctx context.Context, uid uint, productId uint) response.Response {
	resp := response.NewResponse()
	cartDao := dao.NewCartDao(ctx)
	if !cartDao.Exist(uid, productId) {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("cart not exist")
		return resp
	}
	err := cartDao.Delete(uid, productId)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	return resp
}

func (*cartService) UpdateCart(ctx context.Context, productId int, quantity int) response.Response {
	resp := response.NewResponse()

	return resp
}
