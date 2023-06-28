package service

import (
	"context"
	"gin_mall/common"
	"gin_mall/dao"
	"gin_mall/response"
	"gin_mall/vo"
)

// FavoritesService
type favoriteService struct{}

var FavoriteService = favoriteService{}

// AddFavorite
func (*favoriteService) AddFavorite(ctx context.Context, uid, pid uint) response.Response {
	resp := response.NewResponse()
	favoriteDao := dao.NewFavoriteDao(ctx)
	b, err := favoriteDao.IsFavorite(uid, pid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if b {
		resp.SetCode(response.ErrorFavoriteCode)
		resp.SetMessage(response.ErrorFavoriteAlreadyExist)
		return resp
	}
	err = favoriteDao.Create(uid, pid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	return resp
}

// DeleteFavorite
func (*favoriteService) DeleteFavorite(ctx context.Context, uid, pid uint) response.Response {
	resp := response.NewResponse()
	favoriteDao := dao.NewFavoriteDao(ctx)
	b, err := favoriteDao.IsFavorite(uid, pid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if !b {
		resp.SetCode(response.ErrorFavoriteCode)
		resp.SetMessage(response.ErrorFavoriteNotExist)
		return resp
	}
	err = favoriteDao.Delete(uid, pid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	return resp
}

// GetFavoriteList
func (*favoriteService) List(ctx context.Context, uid uint, basePage common.BasePage) response.Response {
	resp := response.NewResponse()
	favoriteDao := dao.NewFavoriteDao(ctx)
	categories, err := favoriteDao.List(uid, basePage.Page, basePage.Size)
	if err != nil {
		resp.SetMessage(err.Error())
		return resp
	}
	resp.Data = vo.ToFavoriteListVo(categories)

	return resp
}
