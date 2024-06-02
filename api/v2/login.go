package v2

import (
	"arlog/middleware"
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
)

// Login 后台登陆
func Login(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	var code int
	//获取账号密码进行验证
	formData, code = service.CheckLogin(formData.Username, formData.Password)
	//登录成功
	if code == 200 {
		//登录成功传入用户名,获取jwt
		token, _ := middleware.GenerateJWT(formData.Username)
		//登录成功返回token和用户信息
		res.Ask(c, code, gin.H{
			"token": token,
			"user":  formData,
		})

		return
	}
	//验证需要otp验证
	if code == 206 {
		res.Ask(c, code, gin.H{
			"opt":  true,
			"user": formData.Username,
		})
		return
	}

	//登录失败,只返回用户名
	res.Ask(c, code, formData.Username)

}
