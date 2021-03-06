package irisapp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"log-api/config"
	"log-api/middleware"
	"log-api/router"
)

var (
	app *iris.Application
)

func InitApp() {
	if nil == app {
		app = iris.New()
		app.Use(middleware.Cors)
		app.Done(middleware.UseAfter)
	}
}

func Start() {
	cfg := config.SrvCfg()
	setup(cfg.Prefix)
	run(cfg.Address)
}

func setup(host string) {
	for key, value := range router.Routers() {
		mvc.Configure(app.Party(host), func (mvcapp *mvc.Application){
			if value.GetNeedAuth() {
				//增加认证鉴权
				mvcapp.Party(key, middleware.Auth).Handle(value.GetController())
			} else {
				//无需认证鉴权
				mvcapp.Party(key).Handle(value.GetController())
			}
		})
	}
}


func run(addr string) {
	app.Run(iris.Addr(addr))
}