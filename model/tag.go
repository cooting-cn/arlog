package model

// Tag 标签表
type Tag struct {
	Id       int    `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	Name     string `gorm:"column:name;type:varchar(50);comment:名称" json:"name"`
	ArtId    int    `gorm:"column:art-id;type:int;comment:关联的文章id" json:"art-id"`
	ModeTime        //添加时间模型
}
