package gorm

import (
	"arlog/core/logs"
	"arlog/utils/res"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
)

var e *casbin.Enforcer

func CasbinInit() {
	text := `
  [request_definition]
  r = sub, obj, act

  [policy_definition]
  p = sub, obj, act

  [policy_effect]
  e = some(where (p.eft == allow))

  [matchers]
  m = r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
  `
	//读取模型
	m, _ := model.NewModelFromString(text)
	//链接数据库
	a, _ := gormadapter.NewAdapterByDB(Db)
	//初始化验证对象
	e, _ = casbin.NewEnforcer(m, a)

	//从数据库加载策略。
	_ = e.LoadPolicy()

	//添加规则到表
	//e.AddPolicy("alice", "data1", "write")

	// 检查权限。
	ok, _ := e.Enforce("admin", "/api/xx", "get")
	logs.Info("验证admin的权限", ok)

	// Save the policy back to DB.
	//e.SavePolicy()
}

// CasbinMiddleWare 权限认证中间件
func CasbinMiddleWare() func(c *gin.Context) {
	return func(c *gin.Context) {
		//从jwt中获取用户名
		username, _ := c.Get("username")

		// 请求的path
		p := c.Request.URL.Path
		// 请求的方法
		m := c.Request.Method
		// 这里认证
		ok, _ := e.Enforce(username, p, m)

		if !ok {
			//返回无权限
			res.Ask(c, 403, nil)

			logs.Info("casbin校验出错无权限", gin.H{
				"ok":       ok,
				"username": username,
				"url":      p,
				"Method":   m,
			})

			c.Abort()
			return
		}
		c.Next()
	}
}
