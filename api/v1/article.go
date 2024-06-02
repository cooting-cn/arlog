package v1

import (
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddArticle 增加文章
func AddArticle(c *gin.Context) {
	//绑定文章结构体
	var formData model.Article
	//绑定结构体json
	_ = c.ShouldBindJSON(&formData)
	//添加文章
	code := service.InsArt(formData)
	//返回结果
	res.Ask(c, code, nil)

}

// UpArticle 修改更新文章
func UpArticle(c *gin.Context) {
	var formData model.Article
	_ = c.ShouldBindJSON(&formData)
	code := service.UpArt(formData)
	res.Ask(c, code, nil)
}

// DelArticle 删除文章
func DelArticle(c *gin.Context) {
	var formData model.Article
	_ = c.ShouldBindJSON(&formData)
	code := service.DelArt(formData.Id)
	res.Ask(c, code, nil)
}

// GetArt 查询所有文章
func GetArt(c *gin.Context) {
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

	formData, code, total := service.GetArt(pageSize, page)
	res.Ask(c, code, gin.H{
		"total": total,
		"data":  formData,
	})

}
