package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
)

// CheckLogin 登录验证
func CheckLogin(username string, password string) (model.User, int) {
	var user model.User
	//查询用户是否存在
	gorm.Db.Where("username = ?", username).First(&user)

	if user.Id == 0 {
		//返回用户不存在
		return user, 204
	}
	//查询是否有otp验证
	if len(user.Otp) > 6 {
		return user, 206
	}
	//对比密码
	ok := ComparePasswords(user.Password, password)
	if !ok {
		//密码错误
		return user, 203
	}
	log.Info("登陆成功,登陆用户信息", user)
	return user, 200

}
