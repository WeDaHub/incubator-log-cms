package controller

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"log-api/common/result"
	"log-api/service"
)

type Accountf struct {

}

// 注册账号
// POST	/account/f/register
func (this *Accountf) PostRegister(ctx iris.Context) string {
	var data struct {
		UserName string 	`json:"account"`
		Password string 	`json:"password"`
		Mobile 	 string 	`json:"mobile"`
		Code	 string 	`json:"code"`
		CodeId	 string		`json:"key"`
	}
	ctx.ReadJSON(&data)
	if len(data.UserName) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Password) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Mobile) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Code) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.CodeId) ==0 {
		return result.CR().Error(1).Json()
	}
	r := service.GetAccountService().AccountRegist(data.UserName, data.Password, data.Mobile, data.Code, data.CodeId)
	return r.Json()
}

// 账号登录
// POST /account/f/login
func (this *Accountf) PostLogin(ctx iris.Context) string {
	var data struct {
		UserName string 	`json:"account"`
		Password string 	`json:"password"`
		Code	 string 	`json:"code"`
		CodeId	 string		`json:"key"`
	}
	err := ctx.ReadJSON(&data)
	fmt.Println(err)
	if len(data.UserName) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Password) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Code) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.CodeId) ==0 {
		return result.CR().Error(1).Json()
	}
	r,token := service.GetAccountService().AccountLogin(data.UserName, data.Password, data.Code, data.CodeId)
	if r.IsSucceed() {
		ctx.Header("Authorization",token)
	}
	return r.Json()
}
