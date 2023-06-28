package dao

import (
	"context"
	"gin_mall/common"
	"gin_mall/model"

	"gorm.io/gorm"
)

type ProductDao struct {
	*gorm.DB
}

func NewProductDao(ctx context.Context) *ProductDao {
	return &ProductDao{
		DB: NewDbCilent(ctx),
	}
}

func (dao *ProductDao) Create(p *model.Product) (err error) {
	return dao.DB.Model(&model.Product{}).Create(p).Error
}

func (dao *ProductDao) CountProductByCondition(condition map[string]interface{}) (total int64, err error) {
	err = dao.DB.Model(&model.Product{}).Where(condition).Count(&total).Error
	return
}

func (dao *ProductDao) ListProductByCondition(condition map[string]interface{},
	page common.BasePage) (products []*model.Product, total int64, err error) {
	info, ok := condition["info"]
	if ok {
		// BUG LIKE 全表扫描？？？？？？？？
		tx := dao.DB.Model(&model.Product{}).Where("title LIKE ? OR info LIKE ?", "%"+info.(string)+"%", "%"+info.(string)+"%")
		tx.Count(&total)
		err = tx.Offset((page.Page - 1) * page.Size).Limit(page.Size).Find(&products).Error
		return
	}

	tx := dao.DB.Model(&model.Product{}).Where(condition)

	tx.Count(&total)
	err = tx.Offset((page.Page - 1) * page.Size).Limit(page.Size).Find(&products).Error

	return
}

// GetProductById
func (dao *ProductDao) GetProductById(id uint) (p *model.Product, err error) {
	err = dao.DB.Model(&model.Product{}).Where(&model.Product{Model: gorm.Model{ID: id}}).First(&p).Error
	return
}

// 获取图片列表
func (dao *ProductDao) ImageList(p *model.Product) []string {
	var images []string
	err := dao.DB.Model(&model.ProductImage{}).Select("image_url").
		Where(&model.ProductImage{ProductId: p.ID}).Find(&images).Error
	if err != nil {
		return nil
	}
	return images
}
