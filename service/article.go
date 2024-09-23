package service

import (
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/model"
)

// InsArt 文章添加
func InsArt(art model.Article) int {

	gorm.Db.Create(&art)

	log.Info("文章写入成功!", art.Id)
	return 200
}

// GetArt 查询文章列表
func GetArt(pageSize int, page int) ([]model.Article, int, int64) {
	var articleList []model.Article

	var total int64
	//分页查询文章列表

	gorm.Db.Offset((page - 1) * pageSize).Limit(pageSize).Order("id desc").Find(&articleList)
	// 统计有多少条数据
	gorm.Db.Model(&articleList).Count(&total)

	log.Info("查询所有文章")
	return articleList, 200, total

}

// DelArt 文章删除
func DelArt(ids []int) int {
	gorm.Db.Delete(&model.Article{}, "id IN ?", ids)
	log.Info("文章删除成功影响行id：", ids)
	return 200
}

// UpArt 文章修改
func UpArt(art model.Article) int {

	gorm.Db.Model(&art).Updates(map[string]interface{}{"id": &art.Id, "title": &art.Title, "sort-id": &art.SortId, "open": &art.Open})

	log.Info("更新成功，受影响的id", art.Id)
	return 200
}

// SearchArt 搜索文章
func SearchArt(keyword string, pageSize, page int) ([]model.Article, int, int64) {
	var articleList []model.Article
	var total int64
	gorm.Db.Offset((page-1)*pageSize).Limit(pageSize).Where("title LIKE?", "%"+keyword+"%").Order("id desc").Find(&articleList).Count(&total)

	log.Info("搜索文章：", keyword, total)
	return articleList, 200, total
}
