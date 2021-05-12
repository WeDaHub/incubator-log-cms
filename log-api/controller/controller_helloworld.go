package controller

import (
	"log-api/common/result"
	"github.com/kataras/iris/v12"
)

type Helloworld struct {

}

func (this*Helloworld) Get(ctx iris.Context) string{
	var data struct {
		Func   string `json:"func"`
	}
	data.Func = "Get"
	return result.CR().Succeed(data).Json()
}

func (this*Helloworld) GetBy(id string, ctx iris.Context) (str string){
	var data struct {
		Func   string `json:"func"`
	}
	data.Func = "GetBy"
	str = result.CR().Succeed(data).Json()
	return
}