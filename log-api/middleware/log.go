package middleware

import (
	"github.com/kataras/iris/v12"
	"log-api/common/log"
)

func UseAfter(ctx iris.Context){
	log.INFO("iris", "%v", ctx)
	ctx.Next()
}
