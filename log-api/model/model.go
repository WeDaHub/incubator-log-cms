package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log-api/config"
	"database/sql"
)
//需要自动迁移的表model添加到此处
var models = []interface{}{
	&User{},
	&Rule{},
	&Content{},
}
//数据库访问对象，建议使用DB()获取
var gormDB *gorm.DB
func DB() *gorm.DB {
	return gormDB
}
//初始化数据库
func Init() (err error) {
	dsn := config.MySQLCfg().GetDSN()
	if gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); nil != err {
		return
	}
	if err = gormDB.AutoMigrate(models...); nil != err {
		return
	}
	var db *sql.DB
	if db, err = gormDB.DB(); nil != err {
		return
	}
	db.SetMaxIdleConns(config.MySQLCfg().MaxIdleCount);
	db.SetMaxOpenConns(config.MySQLCfg().MaxConnCount);
	return nil
}
