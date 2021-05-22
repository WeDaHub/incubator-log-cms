package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/service"
)

type Admin struct {

}

func (this *Admin) Post(ctx iris.Context){
	service.GetAccountService().AccountRegist("leo", "123456", "123")
}