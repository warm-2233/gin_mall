package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/dto"
	"gin_mall/model"
	"gin_mall/response"
	"gin_mall/vo"
)

type addressService struct {
}

var Address = addressService{}

func (*addressService) ListAddresses(ctx context.Context, uid uint) response.Response {
	resp := response.NewResponse()
	addressDao := dao.NewAddressDao(ctx)
	addrs, err := addressDao.List(uid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(response.Error)
		return resp
	}
	var addrs_vo []vo.AddressVo
	for i := 0; i < len(addrs); i++ {
		addrs_vo = append(addrs_vo, vo.AddressVo{
			Aid:     addrs[i].ID,
			Name:    addrs[i].Name,
			Phone:   addrs[i].Phone,
			Address: addrs[i].Address,
		})
	}
	resp.Data = addrs_vo
	return resp
}
func (*addressService) GetAddresses(ctx context.Context, aid uint, uid uint) response.Response {
	resp := response.NewResponse()
	addressDao := dao.NewAddressDao(ctx)
	addr, err := addressDao.Get(aid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if addr.Uid != uid {
		resp.SetCode(response.ErrorUserCode)
		resp.SetMessage("没有权限")
		return resp
	}

	resp.Data = vo.AddressVo{
		Aid:     addr.ID,
		Name:    addr.Name,
		Phone:   addr.Phone,
		Address: addr.Address,
	}

	return resp
}
func (*addressService) AddressCreate(ctx context.Context, uid uint, request dto.AddressCreateDto) response.Response {
	resp := response.NewResponse()
	addressDao := dao.NewAddressDao(ctx)

	addr := &model.Address{
		Uid:     uid,
		Name:    request.Name,
		Phone:   request.Phone,
		Address: request.Address,
	}
	err := addressDao.Create(addr)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.Message = err.Error()
		return resp
	}
	resp.Data = vo.AddressVo{
		Aid:     addr.ID,
		Name:    addr.Name,
		Phone:   addr.Phone,
		Address: addr.Address,
	}

	return resp
}
func (*addressService) AddressUpdate(ctx context.Context, request dto.AddressUpdateDto, uid uint) response.Response {
	resp := response.NewResponse()
	addressDao := dao.NewAddressDao(ctx)
	addr, err := addressDao.Get(request.Aid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if addr.Uid != uid {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("没有权限")
		return resp
	}
	if request.Address != "" {
		addr.Address = request.Address
	}
	if request.Name != "" {
		addr.Name = request.Name
	}
	if request.Phone != "" {
		addr.Phone = request.Phone //TODO: 保险校验电话格式? 可以使用 regexp 或者 特殊
	}

	err = addressDao.Update(addr)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}

	resp.Data = vo.AddressVo{
		Aid:     addr.ID,
		Name:    addr.Name,
		Phone:   addr.Phone,
		Address: addr.Address,
	}
	return resp
}

func (*addressService) AddressDelete(ctx context.Context, aid, uid uint) response.Response {
	resp := response.NewResponse()
	addressDao := dao.NewAddressDao(ctx)
	addr, err := addressDao.Get(aid)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	if addr.Uid != uid {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage("没有权限")
		return resp
	}

	err = addressDao.Delete(addr)
	if err != nil {
		resp.SetCode(response.ErrorCode)
		resp.SetMessage(err.Error())
		return resp
	}
	return resp
}
