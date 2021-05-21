package model

import "gorm.io/gorm"

//角色权限
type Rule struct {
	gorm.Model
	Name	string 	`gorm:"type:varchar(64);not null;comment:'权限名称'"`
	Creator	uint	`gorm:"type:uint;not null;comment:'创建者id;user表主键'"`
}
