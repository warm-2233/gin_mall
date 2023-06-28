package util

import (
	"bytes"
	"errors"
)

// AES 对称加密
type Encryption struct {
	Key string
}

var Encrypt *Encryption

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// PadPwd 填充密码长度
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize

	ret := bytes.Repeat([]byte{byte(padNum)}, padNum)

	return append(srcByte, ret...)
}

// 去掉填充部分
func UnPadPwd(dst []byte) ([]byte, error) {
	length := len(dst)

	if length <= 0 {
		return nil, errors.New("UnPadPwd 长度错误")
	}

	unpadNum := int(dst[length-1])
	if unpadNum > length {
		return nil, errors.New("UnPadPwd 长度错误")
	}

	return dst[:length-unpadNum], nil
}

// set Key
func (e *Encryption) SetKey(key string) {
	e.Key = key
}

// TODO 加解密还没有写啊啊啊啊啊啊啊啊

// 使用AES加密
func (e *Encryption) AesEncoding(s string) string {
	return s
}

// 使用AES解密
func (e *Encryption) AesDecoding(s string) string {
	return s
}
