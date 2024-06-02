package res

import (
	"arlog/utils/er"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Ask 响应请求封装
func Ask(c *gin.Context, code int, data any) {
	c.JSON(200, gin.H{
		"code": code,
		"msg":  er.GetErr(code),
		"data": data,
	})
}
