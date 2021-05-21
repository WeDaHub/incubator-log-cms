package service

import "log-api/common/result"

type account struct {

}


//发送登录验证码
func (this *account) SendLoginSmsCode(mobile string) *result.R{
	return result.CR()
}

//手机验证码登录
func (this *account) MobileLogin(mobile string, smsCode string) (*result.R, string){
	return result.CR(), ""
}
//账号密码登录
func (this *account) AccountLogin(account string, password string) (*result.R,string){
	return result.CR(), ""
}

//账号登出
func (this *account) Logout(userId uint) *result.R{
	return result.CR()
}
//发送找回密码验证码
func (this *account) SendForgetPasswordSmsCode(mobile string) *result.R{
	return result.CR()
}
//发送修改密码验证码
func (this *account) SendModifyPasswordSmsCode(mobile string) *result.R{
	return result.CR()
}

var act *account

func init() {
	act = &account{}
}