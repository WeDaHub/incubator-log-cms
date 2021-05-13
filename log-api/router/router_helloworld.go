package router

import (
	"log-api/controller"
)

type Helloworld struct {
	name string
	contorller interface{}
}

func (this *Helloworld) GetUir() string {
	return this.name
}

func (this *Helloworld) GetController() interface{} {
	return this.contorller
}

func NewHelloworld() Router {
	return &Helloworld{
		name: "/hello",
		contorller: new(controller.Helloworld),
	}
}
