package service

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
	"crypto/rand"
	"log-api/common/result"
)

var store = base64Captcha.DefaultMemStore

type captcha struct {

}

func newDriver() *base64Captcha.DriverString {
	driver := new(base64Captcha.DriverString)
	driver.Height = 44
	driver.Width = 120
	driver.NoiseCount = 5
	driver.ShowLineOptions = base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine
	driver.Length = 6
	driver.Source = "1234567890qwertyuipkjhgfdsazxcvbnm"
	//driver.Fonts = []string{"wqy-microhei.ttc"}
	return driver
}

func calId() string {
	b := make([]byte, 16)
	rand.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}

func (this *captcha) Create() *result.R {
	var data struct {
		Id	 		string 		`json:"id"`
		ImgBase64	string 		`json:"img"`
	}
	data.Id = calId()
	var driver = newDriver().ConvertFonts()
	c := base64Captcha.NewCaptcha(driver, store)
	_, content, answer := c.Driver.GenerateIdQuestionAnswer()
	item, _ := c.Driver.DrawCaptcha(content)
	c.Store.Set(data.Id, answer)
	data.ImgBase64 = item.EncodeB64string()
	return result.CR().Succeed(data)
}
func (this *captcha) Verify(id string, code string) *result.R {
	if store.Verify(id, code, true) {
		return result.CR().Succeed(nil)
	} else {
		return result.CR().ERROR("Verify failed")
	}
}
