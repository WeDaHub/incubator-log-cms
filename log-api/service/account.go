package service

import (
	"log-api/common/result"
	"log-api/middleware"
	"log-api/model"
	"time"
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
	var user model.User
	if err := model.DB().First(&user, "account = ?", account).Error; nil != err {
		return result.CR().ERROR(err.Error()), ""
	}
	if user.ID == 0 {
		return result.CR().ERROR("账号不存在"), ""
	}
	if (user.Password != password){
		return result.CR().ERROR("账号或密码错误"), ""
	}
	//清空密码再生成token，防止token被破解
	user.Password = ""
	if token, err := middleware.GenerateToken(user, time.Hour); nil != err {
		return result.CR().ERROR(err.Error()), ""
	} else {
		return result.CR().Succeed(nil), token
	}
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
