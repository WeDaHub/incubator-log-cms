package model

import "gorm.io/gorm"

type Rule struct {
	gorm.Model
	Name	string 	`gorm:"type:varchar(64);not null;comment:'权限名称'"`
	Creator	uint	`gorm:"type:uint;not null;comment:'创建者id;user表主键'"`
}

//添加角色
func (this *Rule) Add() (err error) {
	err = DB().Create(this).Error
	return
}

//修改角色
func (this *Rule) Modify() (err error) {
	err = DB().Save(this).Error
	return
}

//删除角色信息
func (this *Rule) Delete() (err error) {
	err = DB().Delete(this).Error
	return
}

//通过ID获取角色信息
func (this *Rule) GetById(id uint) (err error) {
	err = DB().Find(this, "id = ?", id).Error
	return
}

func (this *Rule) GetList() (err error) {

	return
}

