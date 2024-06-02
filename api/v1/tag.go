package v1

import (
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddTag 添加Tag
func AddTag(c *gin.Context) {
	var formData model.Tag
	_ = c.ShouldBindJSON(&formData)

	//查询重复是否有Tag
	code := service.GetTag(formData.Name)
	if code == 200 {
		//添加Tag
		code = service.AddTag(formData)
		res.Ask(c, code, formData)
		return
	}
	//重复Tag返回
	res.Ask(c, code, nil)

}

// DelTag 删除Tag
func DelTag(c *gin.Context) {
	var formData model.Tag
	_ = c.ShouldBindJSON(&formData)
	code := service.DelTag(formData)
	res.Ask(c, code, nil)
}

// GetTag 查询所有Tag
func GetTag(c *gin.Context) {

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
		pageSize = 10
	}

	formData, code, total := service.GetTagList(pageSize, page)
	res.Ask(c, code, gin.H{
		"total": total,
		"data":  formData,
	})
}
