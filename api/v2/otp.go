package v2

import (
	"arlog/middleware"
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
)

// Otp 登录开启otp验证
func Otp(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	//从请求链接获取coding编码
	coding, _ := c.GetQuery("coding")
	//开始验证otp
	code, ok := service.OtpToken(coding, formData.Username)
	if ok {
		//登录成功传入用户名,获取jwt
		token, _ := middleware.GenerateJWT(formData.Username)
		//登录成功返回token和用户信息
		res.Ask(c, code, gin.H{
			"user":  formData.Username,
			"token": token,
		})
		return
	}
	//验证失败
	res.Ask(c, code, nil)

}

// BindOtp 绑定otp
func BindOtp(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	//从请求链接获取coding编码
	coding, _ := c.GetQuery("coding")
	//开始绑定otp
	code := service.BindOtp(formData.Username, coding, formData.Otp)
	res.Ask(c, code, nil)
}

// GetOtp 获取otp绑定密钥
func GetOtp(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)
	//开始绑定otp
	url, secret := service.GetOtp(formData.Username)
	res.Ask(c, 200, gin.H{
		"url":    url,
		"secret": secret,
	})
}
