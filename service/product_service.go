package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/dto"
	"gin_mall/model"
	"gin_mall/response"
	"gin_mall/vo"
	"mime/multipart"
	"sync"
)

type productService struct{}

var ProductService = productService{}

func (*productService) Create(ctx context.Context, request dto.ProductDto, uid uint, fs []*multipart.FileHeader) response.Response {
	resp := response.NewResponse()
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetById(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	res := UploadService.UploadToLocalStatic(uid, fs[0], fs[0].Filename)
	if res.Data == nil {
		resp.SetCode(response.ErrorUploadCode)
		resp.SetMessage(err.Error())
		return resp
	}
	data, _ := res.Data.(UploadVo)
	p := data.Path
	product := model.Product{
		Name:          request.Name,
		Category:      request.Category,
		Title:         request.Title,
		Info:          request.Info,
		ImageUrl:      p,
		Price:         request.Price,
		DiscountPrice: request.Discount,
		OnSale:        false,
		Num:           request.Stock,
		BossId:        uid,
		BossName:      user.UserName,
		BossAvatar:    user.Avatar,
	}

	productDao := dao.NewProductDao(ctx)
	err = productDao.Create(&product)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	wg := new(sync.WaitGroup)
	wg.Add(len(fs))
	for i := 0; i < len(fs); i++ {
		productImageDao := dao.NewProductImageDaoByDB(productDao.DB)
		go func(i int) {
			res := UploadService.UploadToLocalStatic(uid, fs[i], "product")
			if res.Data == nil {
				resp.SetCode(response.ErrorCode)
				resp.SetMessage(err.Error())
				return
			}
			data, _ := res.Data.(UploadVo)
			p := data.Path

			productImage := model.ProductImage{
				ProductId: product.ID,
				ImageUrl:  p,
			}
			productImageDao.Create(&productImage)
			wg.Done()
		}(i)
	}
	wg.Wait()

	resp.Data = vo.ToProductVo(&product)
	return resp
}

func (*productService) List(ctx context.Context, request dto.ProductListDto) response.Response {
	resp := response.NewResponse()
	productDao := dao.NewProductDao(ctx)

	condition := make(map[string]interface{})
	if request.Category != 0 {
		condition["categtry_id"] = request.Category
	}
	condition["info"] = request.Info
	// total, err := productDao.CountProductByCondition(condition)
	// if err != nil {
	// 	resp.SetCode(response.ErrorCode)
	// 	resp.SetMessage(err.Error())
	// 	return resp
	// }

	products, total, err := productDao.ListProductByCondition(condition, request.BasePage)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	v := vo.ToProductListVo(products)
	v.Total = int(total)
	resp.Data = v
	return resp
}

// // Search
// func (*productService) Search(ctx context.Context) response.Response {
// 	resp := response.NewResponse()
// 	productDao := dao.NewProductDao(ctx)
// 	condition := make(map[string]interface{})
// 	condition["info"] = service.Info

// 	products, total, err := productDao.ListProductByCondition(condition, service.BasePage)
// 	if err != nil {
// 		resp.SetCode(response.ErrorCode)
// 		resp.SetMessage(err.Error())
// 		return resp
// 	}

// 	resp.Data = gin.H{
// 		"pagination": serializer.ToProductListVo(products),
// 		"total":      total,
// 	}

// 	return resp
// }

// Get Product by id
func GetProductById(ctx context.Context, id uint) response.Response {
	resp := response.NewResponse()
	productDao := dao.NewProductDao(ctx)
	product, err := productDao.GetProductById(id)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = vo.ToProductVo(product)
	return resp
}
