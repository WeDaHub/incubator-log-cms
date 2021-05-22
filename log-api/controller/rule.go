package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/common/anc"
	"log-api/common/result"
	"log-api/service"
)

type Rule struct {

}

// GET /@uri/1
func (this *Rule) GetBy(id int, ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// GET /@uri/list/page/size
func (this *Rule) GetListBy(page int, size int, ctx iris.Context) string {
	return result.CR().Succeed(nil).Json()
}

// POST /@uri
func (this *Rule) Post(ctx iris.Context) string {
	var data struct {
		Level  	int		`json:"level"`
		Name 	string 	`json:"name"`
	}
	ctx.ReadJSON(&data)
	return service.GetRuleService().AddRule(anc.New(ctx), data.Level, data.Name).Json()
}

// PUT /@uri/1
func (this *Rule) PutBy(id int, ctx iris.Context) string {
	var data struct {
		Level  	int		`json:"level"`
		Name 	string 	`json:"name"`
	}
	ctx.ReadJSON(&data)
	return service.GetRuleService().UpdateRule(anc.New(ctx), id, data.Level, data.Name).Json()
}

// DELETE /@uri/1
func (this *Rule) DeleteBy(id int, ctx iris.Context) string {
	return service.GetRuleService().DeleteRule(anc.New(ctx), id).Json()
}

