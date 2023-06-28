package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/response"
	"gin_mall/vo"
)

type categoryService struct {
}

var CategoryService = categoryService{}

func (*categoryService) List(ctx context.Context) response.Response {
	resp := response.NewResponse()
	categoryDao := dao.NewCategoryDao(ctx)
	categories, err := categoryDao.List()
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	resp.Data = categoriesToVO(categories)

	return resp
}

func categoriesToVO(categories []*model.Category) []vo.CategoryVo {
	var vs []vo.CategoryVo
	for _, v := range categories {
		vs = append(vs, vo.CategoryVo{
			CategoryName: v.CategoryName,
			ParentId:     v.ID,
		})
	}
	return vs
}
