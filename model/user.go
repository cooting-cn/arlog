package model

// User 用户表的结构体

type User struct {
	Id       int    `gorm:"column:id;type:int;primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	Username string `gorm:"column:username;type:varchar(100);comment:用户名" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);comment:密码" json:"password"`
	Power    string `gorm:"column:power;type:varchar(10);comment:权限" json:"power"`
	AImg     string `gorm:"column:a-img;type:varchar(255);comment:头像" json:"a-img"`
	Phone    int    `gorm:"column:phone;type:int;comment:手机号" json:"phone"`
	Mail     string `gorm:"column:mail;type:varchar(100);comment:邮箱" json:"mail"`
	Gitee    string `gorm:"column:gitee;type:varchar(100);comment:gitee账号" json:"gitee"`
	Qq       int    `gorm:"column:qq;type:int;comment:qq号" json:"qq"`
	Wechat   string `gorm:"column:wechat;type:varchar(100);comment:微信号" json:"wechat"`
	Ip       string `gorm:"column:ip;type:varchar(255);comment:登陆ip" json:"ip"`
	Otp      string `gorm:"column:otp;type:varchar(255);comment:谷歌令牌" json:"otp"`
	ModeTime        //添加时间模型
}
