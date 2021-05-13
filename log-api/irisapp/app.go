package irisapp

import (
	"log-api/router"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
)

var (
	app *iris.Application
)

func InitApp() {
	if nil == app {
		app = iris.New()
		app.Use(recover.New())
		app.Use(logger.New())
	}
}

func Start(host string, addr string) {
	setup(host)
	run(addr)
}

func setup(host string) {
	for key, value := range router.Routers() {
		mvc.Configure(app.Party(host), func (mvcapp *mvc.Application){
			mvcapp.Party(key).Handle(value.GetController())
		})
	}
}


func run(addr string) {
	app.Run(iris.Addr(addr))
}