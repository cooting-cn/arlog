package v2

import (
	"arlog/middleware"
	"arlog/service"
	"arlog/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
)

// Otp 登录开启otp验证
func Otp(c *gin.Context) {
	var formData struct {
		Username string `json:"username"`
		Coding   string `json:"coding"`
	}
	_ = c.ShouldBindJSON(&formData)

	//开始验证otp
	code, ok, userFormData := service.OtpToken(formData.Coding, formData.Username)
	if ok {
		//登录成功传入用户名,获取jwt
		token, _ := middleware.GenerateJWT(formData.Username)
		userFormData.Password = ""
		//登录成功返回token和用户信息
		res.Ask(c, code, gin.H{
			"token": token,
			"user":  userFormData,
		})
		return
	}
	//验证失败
	res.Ask(c, code, nil)

}

// BindOtp 绑定otp
func BindOtp(c *gin.Context) {
	var formData struct {
		Username string `json:"username"`
		Otp      string `json:"otp"`
		Coding   string `json:"coding"`
	}

	_ = c.ShouldBindJSON(&formData)
	//从请求链接获取coding编码

	//开始绑定otp
	code := service.BindOtp(formData.Username, formData.Coding, formData.Otp)
	res.Ask(c, code, nil)
}

// GetOtp 获取otp绑定密钥
func GetOtp(c *gin.Context) {
	username := c.Query("username")
	fmt.Println(username)
	//开始绑定otp
	url, secret := service.GetOtp(username)
	res.Ask(c, 200, gin.H{
		"url":    url,
		"secret": secret,
	})
}
