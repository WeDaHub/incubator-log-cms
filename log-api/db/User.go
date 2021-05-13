package db

import (
	"fmt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name 	string 	`json:"name"`
	Pwd 	string 	`json:"pwd"`
	Phone 	string 	`json:"phone"`
}

//添加数据
func (user *User) Add() {
	conn := GetDb()
	err := conn.Create(user).Error
	if err != nil {
		fmt.Println("创建失败")
	}
}

//修改数据
func (user *User) Update() {
	conn := GetDb()

	err := conn.Model(user).Updates(user).Error
	if err != nil {
		fmt.Println("修改失败")
	}
}

//删除数据
func (user *User) Del() {
	conn := GetDb()

	err := conn.Delete(user).Error
	if err != nil {
		fmt.Println("删除失败")
	}
}