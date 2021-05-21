package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/result"
)

type Content struct {

}

// PUT /content/a/modify
func (this *Content) PutModify(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// GET /content/a/list
func (this *Content) GetList(ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// GET /content/a/20210521
func (this *Content) GetBy(date string, ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// DELETE /content/a/20210521
func (this *Content) DeleteBy(date string, ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}