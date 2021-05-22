package service

import "log-api/common/result"

type ICaptcha interface {
	Create() *result.R
	Verify(id string, code string) *result.R
}

func GetCaptchaService() ICaptcha {
	return &captcha{}
}
