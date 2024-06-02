package middleware

import (
	"arlog/core"
	"arlog/utils/res"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中

type User struct {
	Username             string `json:"username"`
	jwt.RegisteredClaims        // v5版本新加的方法
}

// GenerateJWT 生成JWT
func GenerateJWT(username string) (string, error) {
	claims := User{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2400 * time.Hour)), // 过期时间24小时
			Issuer:    core.Issuer,                                          // 签发人
		},
	}

	// 使用HS256签名算法
	// 使用指定的签名方法创建签名对象
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	token, err := t.SignedString([]byte(core.CustomSecret))

	return token, err
}

// ParseJwt 解析JWT
func ParseJwt(token string) (*User, error) {
	// 解析token
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	t, err := jwt.ParseWithClaims(token, &User{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(core.CustomSecret), nil
	})

	if err != nil {
		return nil, err
	}

	// 对token对象中的Claim进行类型断言
	if claims, ok := t.Claims.(*User); ok && t.Valid { // 校验token是否有效

		return claims, nil
	} else {

		return nil, errors.New("token验证错误")
	}
}

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中
		// 这里的具体实现方式要依据你的实际业务情况决定
		token := c.Request.Header.Get("token")
		if token == "" {
			//如果为空返回错误
			res.Ask(c, 205, nil)
			c.Abort()
			return
		}

		// 我们使用之前定义好的解析JWT的函数来解析它

		mc, err := ParseJwt(token)
		if err != nil {
			//验证token
			res.Ask(c, 207, err.Error())
			c.Abort()
			return
		}

		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("username", mc.Username)

		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}
