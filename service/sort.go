package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
)

// AddSort 新增分类
func AddSort(sort model.Sort) int {
	err := gorm.Db.Create(&sort).Error
	if err != nil {
		return 201 // 500
	}
	log.Info("添加分类成功", sort)
	return 200
}

// GetSort 查询重复分类
func GetSort(name string) (code int) {
	var sort model.Sort
	gorm.Db.Select("id").Where("name = ?", name).First(&sort)

	if sort.Id > 0 {
		return 211
	}
	return 200
}

// DelSort 删除分类
func DelSort(name string) int {
	var sort model.Sort

	// 根据 name 查找分类
	if err := gorm.Db.Where("name = ?", name).First(&sort).Error; err != nil {
		log.Error("分类未找到:", err)
		return 404 // 返回未找到错误
	}

	// 删除分类
	if err := gorm.Db.Delete(&sort).Error; err != nil {
		log.Error("删除分类失败:", err)
		return 500 // 返回服务器错误
	}

	log.Info("删除分类成功!", sort)
	return 200
}

// GetSortList 查询所有分类
func GetSortList(pageSize int, page int) ([]model.Sort, int, int64) {
	var sortList []model.Sort

	var total int64
	//分页查询分类

	gorm.Db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&sortList)
	// 统计有多少条数据
	gorm.Db.Model(&sortList).Count(&total)

	log.Info("查询所有分类")
	return sortList, 200, total

}
