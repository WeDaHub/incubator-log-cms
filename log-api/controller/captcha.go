package controller

import (
	"github.com/kataras/iris/v12"
	"log-api/service"
)

type Captcha struct {

}

// GET /@uri
func (this *Captcha) Get(ctx iris.Context) string {
	return service.GetCaptchaService().Create().Json()
}

