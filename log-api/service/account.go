package service

import (
	"log-api/common/result"
	"log-api/model"
)

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

//账号注册
func (this *account)AccountRegist(account string, password string, code string) *result.R{
	var user model.User
	if err := model.DB().First(&user, "account = ?", account).Error; nil != err {

		user.Account = account
		user.Password = password

		err = model.DB().Create(&user).Error
		if err != nil {
			return result.CR().Error(1)
		}

	} else {
		return  result.CR().Error(1)
	}

	return result.CR().Succeed(nil)
}

var act *account

func init() {
	act = &account{}
}