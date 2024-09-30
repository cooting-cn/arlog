package v1

import (
	"arlog/model"
	"arlog/service"
	"arlog/utils/res"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// AddSort 添加分类
func AddSort(c *gin.Context) {
	var formData model.Sort
	_ = c.ShouldBindJSON(&formData)

	//查询重复是否有分类
	code := service.GetSort(formData.Name)
	if code == 200 {
		//添加分类
		code = service.AddSort(formData)
		res.Ask(c, code, formData)
		return
	}
	//重复分类返回
	res.Ask(c, code, nil)

}

// DelSort 删除分类
func DelSort(c *gin.Context) {

	var formData struct {
		Name string `json:"name"`
	}

	_ = c.ShouldBindJSON(&formData)
	fmt.Println(formData)
	code := service.DelSort(formData.Name)
	res.Ask(c, code, nil)
}

// GetSort 查询所有分类
func GetSort(c *gin.Context) {

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

	formData, code, total := service.GetSortList(pageSize, page)
	res.Ask(c, code, gin.H{
		"total": total,
		"sorts": formData,
	})
}
