package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

// New NoticeDao
func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{
		NewDbCilent(ctx),
	}
}

// New NoticeDaoByDB
func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{
		db,
	}
}

// GetById 从数据库中获取 Notice
func (dao *NoticeDao) GeNoticeById(uid uint) (Notice *model.Notice, err error) {
	err = dao.DB.Model(model.Notice{}).Where(model.Notice{Model: gorm.Model{ID: uid}}).First(&Notice).Error
	return
}
