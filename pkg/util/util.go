package util

import (
	"gin_mall/conf"
	"os"
	"strconv"
)

func AbsUrl(a string) string {
	return conf.PathConf.Host + conf.PathConf.StaticPath[1:] + a
}

func GetProductUrl(pid uint) string {
	return "http://127.0.0.1:8000/shop/" + strconv.Itoa(int(pid)) + "/"
}

// 判断目录是否存在
func ExistDir(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// 判断文件是否存在
func ExistFile(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
