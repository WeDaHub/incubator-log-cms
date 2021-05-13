package router

import "log-api/controller"

type CommRouter struct {
	uir        string
	controller interface{}
}

func (this *CommRouter) GetUir() string {
	return this.uir
}

func (this *CommRouter) GetController() interface{} {
	return this.controller
}

func newCommRouter(uir string, controller interface{}) Router {
	return &CommRouter{
		uir:        uir,
		controller: controller,
	}
}

func init() {
	//注册通用路由
	register(newCommRouter("/hello", &controller.Helloworld{}))
	register(newCommRouter("/account", &controller.Account{}))
}