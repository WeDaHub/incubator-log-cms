package model

import "gorm.io/gorm"

//角色权限
type Rule struct {
	gorm.Model
	Level 	int 	`gorm:"type:int;not null;comment:'权限等级(数值越大说明权限等级越小)'"`
	Name	string 	`gorm:"type:varchar(64);not null;comment:'权限名称'"`
	Creator	uint	`gorm:"type:uint;not null;comment:'创建者id;user表主键'"`
}
