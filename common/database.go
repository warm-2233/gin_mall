package common

import (
	"context"
	"gin_mall/conf"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"
)

var _db *gorm.DB

func dataDase() {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               conf.MysqlConf.ConnRead(),
		DefaultStringSize: 256,
		// DisableDatetimePrecision: true, // 禁止datetime精度，MySQL 5.6之前数据库不支持
		// DontSupportRenameIndex: true, // 重命名索引，MySQL 5.7之前不支持
		// DontSupportRenameColumn: true, // 用change重命名列， mysql 8 之前不支持
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxIdleConns(10)  // 连接池
	sqlDB.SetMaxOpenConns(100) // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db

	// 主从配置
	_ = _db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{mysql.Open(conf.MysqlConf.ConnWrite())},
		Replicas: []gorm.Dialector{mysql.Open(conf.MysqlConf.ConnRead()),
			mysql.Open(conf.MysqlConf.ConnRead())},
		Policy: dbresolver.RandomPolicy{},
	}))
}

func NewDbCilent(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

// get DB
func GetDb() *gorm.DB {
	return _db
}
