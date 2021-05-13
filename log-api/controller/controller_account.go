package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/result"
)

type Account struct {

}
// 注册账号
// POST	/account/register
func (this *Account) PostRegister(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 账号登录
// POST /account/login
func (this *Account) PostLogin(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 账号登出
// POST /account/logout
func (this *Account) PostLogout(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 修改密码
// PUT /account/password
func (this *Account) PutPassword(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 修改账号信息
// PUT /account/detail
func (this *Account) PutDetail(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 获取账号信息
// GET /account/detail
func (this *Account) GetDetail(ctx iris.Context) string {
	return result.CR().Succeed("GetDetail").Json()
}

// 拉取账号信息列表
// GET /account/details
func (this *Account) GetDetails(ctx iris.Context) string {
	return result.CR().Succeed("GetDetails").Json()
}