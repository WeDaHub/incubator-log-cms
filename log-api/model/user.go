package model

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Mobile 		string 	`gorm:"type:varchar(32);not null;comment:'用户手机号码'"`
	Account		string	`gorm:"type:varchar(64);not null;commnet:'用户账号'"`
	Name 		string 	`gorm:"type:varchar(64);not null;comment:'用户姓名'"`
	Password	string 	`gorm:"type:varchar(64);not null;comment:'登录密码'"`
	RuleId		uint	`gorm:"type:uint;not null;comment:'用户角色权限ID，权限表主键'"`
}

//添加数据
func (user *User) Add() bool {
	if err := DB().Create(user).Error; nil != err {
		fmt.Println(err)
		return false
	}
	return true
}

//修改数据
func (user *User) Update() {
	//if err := DB().Model(user).UpdateColumns()
	//
	//conn := db.GetDb()
	//
	//err := conn.Model(user).Updates(user).Error
	//if err != nil {
	//	fmt.Println("修改失败")
	//}
}

//删除数据
func (user *User) Del() bool {
	if err := DB().Delete(user); nil != err {
		fmt.Println(err)
		return false
	}
	return true
}