package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

//生成otp的绑定密钥

func GetOtp(name string) (url, secret string) {
	// 生成密钥
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "ziit.com",
		AccountName: name,
		Period:      30,
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
	})
	if err != nil {
		log.Panic(err)
	}

	//生成二维码的连接
	log.Info("当前扫码的二维码", key.URL(), key.Secret())
	return key.URL(), key.Secret()
}

// BindOtp 绑定opt令牌
func BindOtp(name, coding, secret string) int {
	// 验证一次性密码
	ok := totp.Validate(coding, secret)
	log.Info("开始绑定", secret, coding, ok)
	if ok {
		log.Info("otp绑定成功", ok)
		//把生成的opt密钥绑定到 用户中
		gorm.Db.Model(&model.User{}).Where("username", name).Update("otp", secret)
		return 200
	}
	log.Info("验证失败", ok)

	return 201

}

// OtpToken 从mysql 获取secret,验证otp令牌
func OtpToken(coding, username string) (int, bool) {
	var user model.User
	//从user表中获取otp的编码
	gorm.Db.Where("username", username).First(&user)

	// 验证一次性密码
	ok := totp.Validate(coding, user.Otp)
	log.Info("开始验证的数据", user.Otp, coding, ok)
	if ok {
		log.Info("验证成功", ok)
		return 200, ok
	} else {
		log.Info("验证失败", ok)
		return 208, ok
	}
}
