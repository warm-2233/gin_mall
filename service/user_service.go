package service

import (
	"context"
	"gin_mall/common"
	"gin_mall/dao"
	"gin_mall/dto"
	"gin_mall/model"
	"gin_mall/pkg/util"
	"gin_mall/response"
	"gin_mall/vo"
	"mime/multipart"
)

type userService struct {
}

var UserService = userService{}

func (*userService) Register(ctx context.Context, request dto.UserRegisterDto) response.Response {
	resp := response.NewResponse()
	userDao := dao.NewUserDao(ctx)
	_, exist, err := userDao.ExistOrNotByUserName(request.UserName)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.Message = err.Error()
		return resp
	}

	if exist {
		resp.SetCode(response.ErrorUserCode)
		resp.Message = "用户名已存在"
		return resp
	}

	user := &model.User{
		Nickname: request.Nickname,
		UserName: request.UserName,
		Status:   common.UserStatusActive,
		Avatar:   "avatar.jpg",
		Money:    10000,
	}

	// 密码加密
	if err := user.SetPassword(request.Password); err != nil {
		resp.SetCode(response.ErrorUserCode)
		resp.SetMessage(response.ErrorFailEncryption)
		return resp
	}

	// 创建用户
	err = userDao.CreateUser(user)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.Message = "创建用户失败"
		return resp
	}

	return resp
}

func (*userService) Login(ctx context.Context, request dto.UserLoginDto) response.Response {
	resp := response.NewResponse()
	var user *model.User
	userDao := dao.NewUserDao(ctx)

	user, exist, err := userDao.ExistOrNotByUserName(request.UserName)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(response.Error)
		return resp
	}

	if !exist {
		resp.SetCode(response.ErrorUserCode)
		resp.Message = response.ErrorNotExistUser
		return resp
	}

	if !user.CheckPassword(request.Password) {
		// 密码错误
		resp.SetCode(response.ErrorUserCode)
		resp.Message = response.ErrorWrongUsernameOrPassword
		return resp
	}

	// token
	token, err := util.GenerateToken(user.ID, user.UserName, user.Authority)
	if err != nil {
		resp.SetCode(response.ErrorAuthCode)
		resp.Message = "token 验证失败"
		return resp
	}

	resp.SetData(vo.LoginUserVo{
		Token: token,
		User:  *vo.UserToVo(user),
	})

	return resp
}

// 更新
func (*userService) Update(ctx context.Context, uid uint, request dto.UserUpdateDto) response.Response {
	resp := response.NewResponse()
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetById(uid)
	if err != nil || user.ID == 0 || user.ID != uid {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(response.ErrorNotPermission)
		return resp
	}

	// 修改昵称
	if request.Nickname != "" {
		user.Nickname = request.Nickname
	}
	if request.UserName != "" {
		u, exist, err := userDao.ExistOrNotByUserName(request.UserName)
		if err != nil {
			resp.SetCode(response.ErrorCode)
			resp.Message = err.Error()
			return resp
		}

		if exist && u.ID != user.ID {
			resp.SetCode(response.ErrorUserCode)
			resp.Message = "用户名已存在"
			return resp
		}
		user.UserName = request.UserName
	}
	if request.Email != "" { // BUG 邮箱验证
		user.Email = request.Email
	}
	if request.Avatar != "" {
		user.Avatar = request.Avatar
	}

	// 更新用户信息
	err = userDao.UserUpdate(user)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("更新用户信息失败")
		return resp
	}

	resp.Data = vo.UserToVo(user)

	return resp
}

// 上传头像
func (*userService) UploadAvatar(ctx context.Context, uid uint, file *multipart.FileHeader) response.Response {
	resp := response.NewResponse()

	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetById(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	// 保存图片到本地
	res := UploadService.UploadToLocalStatic(uid, file, "avatar")
	if res.Data == nil {
		resp.SetCode(response.ErrorUploadCode)
		resp.Message = "上传头像失败"
		return resp
	}

	data, _ := res.Data.(UploadVo)
	p := data.Path
	// 修改用户头像
	user.Avatar = p
	err = userDao.UserUpdate(user)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = vo.UserToVo(user)

	return resp
}

func (*userService) ShowMoney(ctx context.Context, uid uint) response.Response {
	resp := response.NewResponse()
	userDao := dao.NewUserDao(ctx)
	user, err := userDao.GetById(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = vo.MoneyVo{Money: user.Money}

	return resp
}
