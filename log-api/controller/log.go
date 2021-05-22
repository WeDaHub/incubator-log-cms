package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/anc"
	"log-api/service"
)

type Log struct {

}

// PUT /@uri/20210521
func (this *Log) PutBy(date string, ctx iris.Context) string {
	var data struct {
		Content 	string 	`json:"content"`
	}
	ctx.ReadJSON(&data)
	return service.GetLogService().UpdateLog(anc.New(ctx), date, data.Content).Json()
}

// GET /@uri/list/page/size
func (this *Log) GetListBy(page int, size int, ctx iris.Context) string {
	defer ctx.Next()
	return service.GetLogService().GetLogList(anc.New(ctx), page, size).Json()
}

// GET /@uri/20210521
func (this *Log) GetBy(date string, ctx iris.Context) string {
	return service.GetLogService().GetLog(anc.New(ctx), date).Json()
}

// DELETE /@uri/20210521
func (this *Log) DeleteBy(date string, ctx iris.Context) string {
	return service.GetLogService().DeleteLog(anc.New(ctx), date).Json()
}
