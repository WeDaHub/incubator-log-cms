package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/result"
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
	return result.CR().Succeed(nil).Json()
}

func (this *Accountf) GetTest(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}