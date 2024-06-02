package model

// Sort 分类表结构
type Sort struct {
	Id       int    `gorm:"column:id;type:int;comment:id;NOT NULL" json:"id"`
	Name     string `gorm:"column:name;type:varchar(50);comment:分类" json:"name"`
	ModeTime        //添加时间模型
}
