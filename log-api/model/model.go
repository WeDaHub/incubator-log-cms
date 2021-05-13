package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log-api/config"
)
//需要自动迁移的表model添加到此处
var models = []interface{}{
	&User{},
}
//数据库访问对象，建议使用DB()获取
var db *gorm.DB
func DB() *gorm.DB {
	return db
}
//初始化数据库
func Init() (err error) {
	dsn := config.MySQLCfg().GetDSN()
	if db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); nil != err {
		return
	}
	if err = db.AutoMigrate(models...); nil != err {
		return
	}
	return nil
}
