package model

import "gorm.io/gorm"

// 日志内容
type Log struct {
	gorm.Model
	UserId 		uint	`gorm:"type:uint;not null;comment:'用户ID，用户表主键'"`
	DateTime 	string 	`gorm:"type:varchar(32);not null;comment:'日志日期20210521'"`
	Content		string	`gorm:"type:text;not null;comment:'日志内容'"`
}
