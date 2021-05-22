package middleware

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"net/http"
)

func Auth(ctx iris.Context) {
	fmt.Println("auth")
	// 通过jwt 验证是否已登录鉴权
	if jwtAuth(ctx) {
		//已登录成功
		ctx.Next()
	} else {
		//未登录无法使用此接口
		ctx.StatusCode(http.StatusUnauthorized)
		ctx.EndRequest()
	}
}

func jwtAuth(ctx iris.Context) bool {
	token := ctx.GetHeader("Authorization")

	if _, err := VerifyToken(token); nil != err {
		return false
	}
	return true
}
