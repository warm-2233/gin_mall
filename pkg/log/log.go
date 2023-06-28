package log

import (
	"gin_mall/pkg/util"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	if Logger == nil {
		Logger = logrus.New()
	}

	f := setOutputFile()
	Logger.Out = f
	Logger.SetLevel(logrus.DebugLevel)
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})
}

// 日志输出文件
func setOutputFile() *os.File {
	now := time.Now()

	// 存放log路径
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	logPath := wd + "/logs/"

	// 创建日志目录
	if !util.ExistDir(logPath) {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// 创建日志文件
	fileName := logPath + now.Format("2006-01-02_15-04-05") + ".log"

	// 创建日志文件
	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

func InfoLn(args ...interface{}) {
	Logger.Infoln(args...)
}
