package routes

import (
	"arlog/api"
	v1 "arlog/api/v1"
	v2 "arlog/api/v2"
	"arlog/core"
	"arlog/core/gorm"
	log "arlog/core/logs"
	"arlog/middleware"
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
)

func WebInit() {
	//取消gin的模式
	//gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(log.LoggerToFile(), gin.Recovery())
	r.Use(middleware.Cors()) //添加跨域
	// 服务检查
	r.GET("/version", func(c *gin.Context) {
		//返回正常
		res.Ask(c, 200, "version 正常")

	})

	//admin后端接口  添加jwt认证  ,添加权限校验
	admin := r.Group("api/v1").Use(middleware.JWTAuthMiddleware(), gorm.CasbinMiddleWare())
	{
		admin.GET("/", func(c *gin.Context) {

			res.Ask(c, 200, "api/v1正常")

		})

		admin.GET("get-otp", api.GetOtp)    //获取otp
		admin.POST("bind-otp", api.BindOtp) //绑定otp

		admin.POST("add-user", v1.AddUser)       //添加用户
		admin.POST("del-user", v1.DelUser)       //删除用户
		admin.GET("get-all-user", v1.GetAllUser) //查询所有用户
		admin.POST("up-user", v1.UpUser)         //更新资料

		admin.POST("add-art", v1.AddArticle) //添加文章
		admin.GET("get-art", v1.GetArt)      //查询所有文章
		admin.POST("del-art", v1.DelArticle) //删除文章
		admin.POST("up-art", v1.UpArticle)   //更新文章

		admin.POST("add-sort", v1.AddSort) //添加分类
		admin.POST("del-sort", v1.DelSort) //删除分类
		//admin.POST("up-sort", v1.UpSort)//更新分类
		admin.GET("get-sort", v1.GetSort) //查询分类

		admin.POST("add-tag", v1.AddTag) //添加tag
		admin.POST("del-tag", v1.DelTag) //删除tag
		//admin.POST("up-tag", v1.UpTag)//更新tag
		admin.GET("get-tag", v1.GetTag) //查询tag

		admin.POST("up-file", v1.UpFile) //上传文件
		//登录模块
		//文章展示模块
		//分类展示导航模块
		//云标签模块
		//博客运行信息模块
		//热点文章模块
		//三方登录模块

	}
	//v2首页展示接口
	home := r.Group("api/v2")
	{
		home.GET("/", func(c *gin.Context) {
			res.Ask(c, 200, "api/v2正常")
		})
		//用户登录

		home.POST("login", v2.Login) //三方登录code处理
		home.POST("otp", api.Otp)    //otp 验证登录
		//搜索模块
		//文章模块
		//分类模块
		//评价模块
		//文件推荐模块

	}

	err := r.Run(":" + core.WebPort) //运行端口
	log.Error(err)

}
