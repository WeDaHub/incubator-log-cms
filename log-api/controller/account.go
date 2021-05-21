package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/result"
)

type Account struct {

}

// 账号登出
// POST /account/a/logout
func (this *Account) PostLogout(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 修改密码
// PUT /account/a/password
func (this *Account) PutPassword(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 修改账号信息
// PUT /account/a/detail
func (this *Account) PutDetail(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// 获取账号信息
// GET /account/a/detail
func (this *Account) GetDetail(ctx iris.Context) string {
	return result.CR().Succeed("GetDetail").Json()
}

// 拉取账号信息列表
// GET /account/a/details
func (this *Account) GetDetails(ctx iris.Context) string {
	return result.CR().Succeed("GetDetails").Json()
}
