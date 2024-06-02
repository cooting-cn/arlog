package v1

import (
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
)

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	//判断用户是否存在
	code := service.QueUser(formData)
	if code == 202 {
		res.Ask(c, code, nil)
		return
	}

	//添加成功
	formData, code = service.AddUser(formData)
	res.Ask(c, code, formData)
}

// GetAllUser 查询所有用户
func GetAllUser(c *gin.Context) {

	res.Ask(c, 200, service.AllUser())
}

// UpUser 更新用户信息
func UpUser(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	code := service.UpUser(formData)
	res.Ask(c, code, formData)
}

// DelUser 删除用户
func DelUser(c *gin.Context) {
	var formData model.User
	_ = c.ShouldBindJSON(&formData)

	code := service.DelUser(formData)
	res.Ask(c, code, nil)
}
