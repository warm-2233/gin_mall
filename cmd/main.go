package main

import (
	"gin_mall/common"
	"gin_mall/conf"
	"gin_mall/dao"
	"gin_mall/routes"
)

func init() {
	conf.Setup()
	common.Setup()
	dao.Migration()
}

// @title		mall 后台API服务
// @version	1.0
func main() {
	r := routes.NewRouter()

	err := r.Run(conf.ServerConf.HttpPort)
	if err != nil {
		panic(err)
	}
}
