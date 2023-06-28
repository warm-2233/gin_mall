package service

import (
	"fmt"
	"gin_mall/conf"
	"gin_mall/pkg/util"
	"gin_mall/response"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

type uploadService struct{}

var UploadService = uploadService{}

func (*uploadService) UploadToLocalStatic(uid uint, file *multipart.FileHeader, basePath string) response.Response {
	resp := response.NewResponse()
	staticPath := path.Join(basePath, fmt.Sprintf("%d_%s", uid, file.Filename))
	filePath := path.Join(conf.PathConf.StaticPath, staticPath)
	f, err := file.Open()
	if err != nil {
		resp.SetCode(response.ErrorUploadCode)
		resp.SetMessage("图片上传失败" + err.Error())
		return resp
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		resp.SetCode(response.ErrorUploadCode)
		resp.SetMessage("图片上传失败2" + err.Error())
		return resp
	}
	if !DirExistOrNot(path.Dir(filePath)) {
		CreateDir(path.Dir(filePath))
	}
	err = ioutil.WriteFile(filePath, content, 0644)
	if err != nil {
		resp.SetCode(response.ErrorUploadCode)
		resp.SetMessage("图片上传失败3" + err.Error())
		return resp
	}

	resp.SetData(UploadVo{staticPath, util.AbsUrl(staticPath)})
	return resp
}

type UploadVo struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func DirExistOrNot(p string) bool {
	s, err := os.Stat(p)
	if err != nil {
		if os.IsExist(err) {
			goto OK
		}
		return false
	}
OK:
	return s.IsDir()
}

// 创建文件夹
func CreateDir(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}
