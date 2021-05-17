package router

import "log-api/controller"

type CommRouter struct {
	uir        	string
	controller 	interface{}
	needAuth	bool
}

func (this *CommRouter) GetUir() string {
	return this.uir
}

func (this *CommRouter) GetController() interface{} {
	return this.controller
}

func (this *CommRouter) GetNeedAuth() bool {
	return this.needAuth
}

func newCommRouter(uir string, controller interface{}, needAuth bool) Router {
	return &CommRouter{
		uir:        uir,
		controller: controller,
		needAuth: needAuth,
	}
}

func newCommRouterF(uir string, controller interface{}) Router {
	return &CommRouter{
		uir:        uir + "/f",
		controller: controller,
		needAuth: false,
	}
}

func newCommRouterA(uir string, controller interface{}) Router {
	return &CommRouter{
		uir:        uir + "/a",
		controller: controller,
		needAuth: true,
	}
}

func registerAFRouter(uir string, a interface{}, f interface{}) {
	register(newCommRouterA(uir, a))
	register(newCommRouterF(uir, f))
}

func init() {
	//注册通用路由
	register(newCommRouter("/hello", &controller.Helloworld{}, false))
	registerAFRouter("/account", &controller.Account{}, &controller.Accountf{})
}
