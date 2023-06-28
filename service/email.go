package service

import (
	"context"
	"fmt"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/util"
	"gin_mall/response"
	"gin_mall/vo"
	"strings"
	"time"

	"gopkg.in/mail.v2"
)

type SendEmailService struct {
	Email         string `json:"email" form:"email"`
	Password      string `json:"password" form:"password"`
	OperationType uint   `json:"operation_type" form:"operation_type"`
}

const (
	//NOTE OperationType 对应数据库中三条数据，值为数据ID
	OperationTypeBindEmail   = 1
	OperationTypeUnbindEmail = 2
	OperationTypeUpdatePWD   = 3
)

// SendEmail
func (service *SendEmailService) SendEmail(ctx context.Context, uid uint) response.Response {
	var notive *model.Notice
	resp := response.NewResponse()
	token, err := util.GenerateEmailToken(uid, service.OperationType,
		service.Email, service.Password)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		return resp
	}

	notiveDao := dao.NewNoticeDao(ctx)
	notive, err = notiveDao.GeNoticeById(service.OperationType)
	if err != nil {
		resp.SetCode(response.ErrorAuthCode)
		resp.SetMessage(response.ErrorAuthToken)
		return resp
	}
	address := conf.EmailConf.ValidEmail + token
	mailStr := notive.Text
	mailTex := strings.Replace(mailStr, "{{address}}", address, -1)

	m := mail.NewMessage()
	m.SetHeader("From", conf.EmailConf.SmtpEmail)
	m.SetHeader("To", service.Email)
	m.SetHeader("Subject", "您的邮箱验证")

	m.SetBody("text/html", mailTex)

	fmt.Printf("mailTex: %v\n", mailTex)
	fmt.Printf("address: %v\n", address)

	d := mail.NewDialer(conf.EmailConf.SmtpHost, conf.EmailConf.SmtpPort, conf.EmailConf.SmtpEmail, conf.EmailConf.SmtpPwd)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	err = d.DialAndSend(m)
	if err != nil {
		resp.SetCode(response.ErrorEmailCode)
		resp.SetMessage(err.Error())
		return resp
	}

	return resp
}

type ValidateEmailService struct {
}

// ValidateEmail
// 解析流程
// 1. 登录发送邮件
// 2. 邮件中获得token
// 3. 解析token 执行对应操作
func (service *ValidateEmailService) ValidateEmail(ctx context.Context, token string) response.Response {
	resp := response.NewResponse()
	var uid uint
	var email string
	var password string
	var operationType uint

	// 验证 token
	if token == "" {
		resp.SetCode(response.ErrorAuthCode)
		resp.SetMessage(response.ErrorAuthToken)
		return resp
	}

	// 解析token
	claims, err := util.ParseEmailToken(token)
	if err != nil {
		resp.SetCode(response.ErrorAuthCode)
		resp.SetMessage(err.Error())
		return resp
	}

	// 检查token是否过期
	if claims.ExpiresAt < time.Now().Unix() {
		resp.SetCode(response.ErrorAuthCode)
		resp.SetMessage(response.ErrorAuthCheckTokenTimeout)
		return resp
	}
	uid = claims.ID
	email = claims.Email
	password = claims.Password
	operationType = claims.OperationType

	// 获取该用户信息
	dao := dao.NewUserDao(ctx)
	user, err := dao.GetById(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	switch operationType {
	case OperationTypeBindEmail:
		user.Email = email
	case OperationTypeUnbindEmail:
		user.Email = ""
	case OperationTypeUpdatePWD:
		err := user.SetPassword(password)
		if err != nil {
			resp.SetCode(response.ErrorCode)
			resp.SetMessage(err.Error())
			return resp
		}
	}

	err = dao.UserUpdate(user)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = vo.UserToVo(user)
	return resp
}
