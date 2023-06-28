package conf

import (
	"fmt"
	"strconv"

	"gopkg.in/ini.v1"
)

type serverConf struct {
	AppMode  string
	HttpPort string
}

type mysqlConf struct {
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
}

type redisConf struct {
	RedisDb     int
	RedisAddr   string
	RedisPwd    string
	RedisDbName string
}

type qiniuConf struct {
	AccessKey   string
	SecretKey   string
	Bucket      string
	QiniuServer string
}

type emailConf struct {
	ValidEmail string
	SmtpHost   string
	SmtpPort   int
	SmtpEmail  string
	SmtpPwd    string
}

type pathConf struct {
	Host       string
	StaticPath string
}

var (
	MysqlConf  = mysqlConf{}
	ServerConf = serverConf{}
	RedisConf  = redisConf{}
	QiniuConf  = qiniuConf{}
	EmailConf  = emailConf{}
	PathConf   = pathConf{}
)

func Setup() {
	// 读取本地环境变量
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
	}

	loadServer(file)
	loadMySql(file)
	loadRedis(file)
	loadEmail(file)
	loadPhotoPath(file)
	loadQiniu(file)
}

func loadServer(f *ini.File) {
	ServerConf.AppMode = f.Section("service").Key("AppMode").String()
	ServerConf.HttpPort = f.Section("service").Key("HttpPort").String()
}

func loadMySql(f *ini.File) {
	MysqlConf.Db = f.Section("mysql").Key("Db").String()
	MysqlConf.DbHost = f.Section("mysql").Key("DbHost").String()
	MysqlConf.DbPort = f.Section("mysql").Key("DbPort").String()
	MysqlConf.DbUser = f.Section("mysql").Key("DbUser").String()
	MysqlConf.DbPassword = f.Section("mysql").Key("DbPassword").String()
	MysqlConf.DbName = f.Section("mysql").Key("DbName").String()
}
func loadRedis(f *ini.File) {
	i, _ := strconv.Atoi(f.Section("redis").Key("RedisDb").String())
	RedisConf.RedisDb = i
	RedisConf.RedisAddr = f.Section("redis").Key("RedisAddr").String()
	RedisConf.RedisPwd = f.Section("redis").Key("RedisPwd").String()
	RedisConf.RedisDbName = f.Section("redis").Key("RedisDbName").String()
}
func loadEmail(f *ini.File) {
	EmailConf.ValidEmail = f.Section("email").Key("ValidEmail").String()
	EmailConf.SmtpHost = f.Section("email").Key("SmtpHost").String()
	EmailConf.SmtpEmail = f.Section("email").Key("SmtpEmail").String()
	EmailConf.SmtpPwd = f.Section("email").Key("SmtpPwd").String()
	i, _ := strconv.Atoi(f.Section("email").Key("SmtpPort").String())
	EmailConf.SmtpPort = i
}
func loadPhotoPath(f *ini.File) {
	PathConf.Host = f.Section("path").Key("Host").String()
	PathConf.StaticPath = f.Section("path").Key("StaticPath").String()
}

func loadQiniu(f *ini.File) {
	QiniuConf.AccessKey = f.Section("qiniu").Key("AccessKey").String()
	QiniuConf.SecretKey = f.Section("qiniu").Key("SecretKey").String()
	QiniuConf.Bucket = f.Section("qiniu").Key("Bucket").String()
	QiniuConf.QiniuServer = f.Section("qiniu").Key("QiniuServer").String()
}

// 读数据的MySQL路径
func (conf *mysqlConf) ConnRead() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
		"charset=utf8mb4&parseTime=true",
	)
}

// 写数据的MySQL路径
func (conf *mysqlConf) ConnWrite() string {
	return conf.ConnRead()
}
