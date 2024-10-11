package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
)

// AddTag 新增tag
func AddTag(tag model.Tag) int {
	err := gorm.Db.Create(&tag).Error
	if err != nil {
		return 201 // 500
	}
	log.Info("添加tag成功", tag)
	return 200
}

// GetTag 查询重复tag
func GetTag(name string) (code int) {
	var tag model.Tag
	gorm.Db.Select("id").Where("name = ?", name).First(&tag)

	if tag.Id > 0 {
		return 215
	}
	return 200
}

// DelTag 删除tag
func DelTag(name string) int {

	var tag model.Tag

	// 根据 name 查找tag
	if err := gorm.Db.Where("name = ?", name).First(&tag).Error; err != nil {
		log.Error("分类未找到:", err)
		return 404 // 返回未找到错误
	}

	// 删除tag
	if err := gorm.Db.Delete(&tag).Error; err != nil {
		log.Error("删除分类失败:", err)
		return 500 // 返回服务器错误
	}

	log.Info("删除分类成功!", tag)
	return 200

}

// GetTagList 查询所有tag
func GetTagList(pageSize int, page int) ([]model.Tag, int, int64) {
	var tagList []model.Tag

	var total int64
	//分页查询分类

	gorm.Db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&tagList)
	// 统计有多少条数据
	gorm.Db.Model(&tagList).Count(&total)

	log.Info("查询所有tag")
	return tagList, 200, total

}
