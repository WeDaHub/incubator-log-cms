package model

import (
	"gorm.io/gorm"
)

//用户表
type User struct {
	gorm.Model
	Mobile 		string 	`gorm:"type:varchar(32);not null;comment:'用户手机号码'"`
	Account		string	`gorm:"type:varchar(64);not null;commnet:'用户账号'"`
	Name 		string 	`gorm:"type:varchar(64);not null;comment:'用户姓名'"`
	Password	string 	`gorm:"type:varchar(64);not null;comment:'登录密码'"`
	RuleId		uint	`gorm:"type:uint;not null;comment:'用户角色权限ID，权限表主键'"`
}
