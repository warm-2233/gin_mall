package dao

import (
	"context"
	"gin_mall/model"

	"gorm.io/gorm"
)

type UserDao struct {
	*gorm.DB
}

// New UserDao
func NewUserDao(ctx context.Context) *UserDao {
	return &UserDao{
		NewDbCilent(ctx),
	}
}

// New UserDaoByDB
func NewUserDaoByDB(db *gorm.DB) *UserDao {
	return &UserDao{
		db,
	}
}

// ExistOrNotByUserName 根据用户名查询是否存在
func (dao *UserDao) ExistOrNotByUserName(userName string) (user *model.User, exist bool, err error) {
	err = dao.DB.Model(&model.User{}).Where(model.User{UserName: userName}).First(&user).Error

	// 没有找到
	if user == nil || err == gorm.ErrRecordNotFound {
		return nil, false, nil
	}
	// 找到了
	return user, true, err
}

// CreateUser 创建用户
func (dao *UserDao) CreateUser(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Create(user).Error
	return
}

// GetById 从数据库中获取用户
func (dao *UserDao) GetById(uid uint) (user *model.User, err error) {
	err = dao.DB.Model(model.User{}).Where(model.User{Model: gorm.Model{ID: uid}}).First(&user).Error
	return
}

// UserUpdate 修改用户
func (dao *UserDao) UserUpdate(user *model.User) (err error) {
	err = dao.DB.Model(&model.User{}).Where(model.User{Model: gorm.Model{ID: user.ID}}).Updates(user).Error
	return
}
