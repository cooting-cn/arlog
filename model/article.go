package model

// Article 对应文字表
type Article struct {
	Id       int    `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	Title    string `gorm:"column:title;type:varchar(100);comment:标题" json:"title"`
	Desc     string `gorm:"column:desc;type:longtext;comment:描述" json:"desc"`
	Content  string `gorm:"column:content;type:longtext;comment:内容" json:"content"`
	Img      string `gorm:"column:img;type:longtext;comment:图片" json:"img"`
	Open     uint   `gorm:"column:open;type:int unsigned;comment:是否公开;NOT NULL" json:"open"`
	SortId   int    `gorm:"column:sort-id;type:int;comment:分类" json:"sortId"`
	ModeTime        //添加时间模型
}
