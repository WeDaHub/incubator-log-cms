package anc

import (
	"github.com/kataras/iris/v12"
	"log-api/middleware"
	"log-api/model"
)

type Context struct {
	Ctx			iris.Context
	User 		model.User
	IsLogined	bool
}

func New(ctx iris.Context) Context {
	token := ctx.GetHeader("Authorization")
	claims, err := middleware.VerifyToken(token)
	if nil != err {
		return Context{
			Ctx: ctx,
			User: model.User{},
			IsLogined: false,
		}
	} else {
		return Context{
			Ctx: ctx,
			User: claims.User,
			IsLogined: true,
		}
	}
}