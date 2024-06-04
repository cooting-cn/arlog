package main

import (
	"arlog/core"
	"arlog/core/gorm"
	"arlog/core/logs"
	"arlog/core/object"
	"arlog/routes"
)

func main() {
	logs.InitLog()    //初始化log
	core.ConfInit()   //初始化配置
	gorm.DbInit()     //初始化mysql
	gorm.CasbinInit() //初始化权限
	object.CosInit()  //初始化对象存储
	routes.WebInit()  //初始化web
}
