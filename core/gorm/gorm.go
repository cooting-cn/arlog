package gorm

import (
	"arlog/core"
	log "arlog/core/logs"
	"arlog/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB //使用全局变量
var err error   //使用全局变量
func DbInit() {

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		core.MysqlName, core.MysqlPassword, core.MysqlHost, core.MysqlPort, core.MysqlDb)

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		log.Panic(err)
		return
	}

	//配置sql连接池
	dbSql, _ := Db.DB()
	err = dbSql.Ping()
	if err != nil {
		log.Panic(err)
		return
	}
	dbSql.SetMaxIdleConns(10)               //空闲连接数10
	dbSql.SetMaxOpenConns(100)              //数据库最大连接数100
	dbSql.SetConnMaxLifetime(1 * time.Hour) //最大复用时间10秒
	// 迁移 注册表
	_ = Db.AutoMigrate(&model.User{}, &model.Article{}, &model.Sort{}, &model.Tag{})

	//初始化admin用户,如果存在就跳过
	var user model.User
	Db.First(&user)
	if user.Username == "admin" {
		log.Info("查询到admin用户, 停止添加")
	} else {
		//创建添加密码
		user = model.User{Username: "admin", Password: "$2y$04$eq9KPY9gp8uZa9rc65qW7eWGmkaqc4chU1H1dy/5fB/7TyXVmpSFO"}
		Db.Create(&user)
		log.Info("没有admin用户,添加admin")

	}
	//Db.Create(&user)

}
