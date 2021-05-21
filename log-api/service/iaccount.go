package service

import (
	"log-api/common/result"
)

type IAccount interface {
	//发送登录验证码
	SendLoginSmsCode(mobile string) *result.R
	//手机验证码登录
	MobileLogin(mobile string, smsCode string) (*result.R,string)
	//账号密码登录
	AccountLogin(account string, password string) (*result.R, string)
	//账号登出
	Logout(userId uint) *result.R
	//发送找回密码验证码
	SendForgetPasswordSmsCode(mobile string) *result.R
	//发送修改密码验证码
	SendModifyPasswordSmsCode(mobile string) *result.R
	//账号注册
	AccountRegist(account string, password string, code string) *result.R
}

func GetAccountService() IAccount {
	return act
}