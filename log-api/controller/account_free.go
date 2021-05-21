package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/result"
	"log-api/service"
)

type Accountf struct {

}

// 注册账号
// POST	/account/f/register
func (this *Accountf) PostRegister(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 账号登录
// POST /account/f/login
func (this *Accountf) PostLogin(ctx iris.Context) string {
	var data struct {
		UserName string 	`json:"account"`
		Password string 	`json:"password"`
	}
	ctx.ReadJSON(data)
	if len(data.UserName) ==0 {
		return result.CR().Error(1).Json()
	}
	if len(data.Password) ==0 {
		return result.CR().Error(1).Json()
	}
	r,token := service.GetAccountService().AccountLogin(data.UserName, data.Password)
	if r.IsSucceed() {
		ctx.Header("Authorization",token)
	}
	return r.Json()
}
