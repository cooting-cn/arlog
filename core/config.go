package core

import (
	log "arlog/core/logs"
	"github.com/spf13/viper"
)

var (
	WebPort       string
	MysqlHost     string
	MysqlPort     string
	MysqlName     string
	MysqlPassword string
	MysqlDb       string
	Url           string
	SecretID      string
	SecretKey     string
	Path          string
	CustomSecret  string
	Issuer        string
)

// ConfInit 获取本地配置文件
func ConfInit() {
	viper.SetConfigName("config.json") //配置文件名
	viper.SetConfigType("json")        // 配置文件类型
	viper.AddConfigPath("./")          // 配置文件路径
	err := viper.ReadInConfig()        // 读取配置文件信息
	if err != nil {
		log.Panic("config配置加载错误", err)
	}
	WebPort = viper.GetString("web.port")              //获取web运行端口
	MysqlHost = viper.GetString("mysql.host")          //获取配置mysql的连接ip
	MysqlPort = viper.GetString("mysql.port")          //获取mysql的端口
	MysqlName = viper.GetString("mysql.name")          //获取mysql的用户名
	MysqlPassword = viper.GetString("mysql.password")  //获取mysql的密码
	MysqlDb = viper.GetString("mysql.db")              //获取mysql运行的库
	Url = viper.GetString("object.url")                //对象存储的访问地址
	SecretID = viper.GetString("object.id")            //对象存储的id
	SecretKey = viper.GetString("object.key")          //对象存储的key
	Path = viper.GetString("object.path")              //存储路径
	CustomSecret = viper.GetString("jwt.CustomSecret") //jwt-CustomSecret
	Issuer = viper.GetString("jwt.Issuer")             //获取jwt-Issuer
}
