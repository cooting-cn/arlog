package v1

import (
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
	"strconv"
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
	//页数
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	//行数
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 12
	}
	formData, code, total := service.AllUser(pageSize, page)
	res.Ask(c, code, gin.H{
		"total": total,
		"tags":  formData,
	})

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
