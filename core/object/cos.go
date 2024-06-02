package object

import (
	"arlog/core"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
)

var Cos *cos.Client

// CosInit 初始化对象存储
func CosInit() {
	//创建cos客户端并设置SecretID，SecretKey和访问域名
	u, _ := url.Parse(core.Url)
	b := &cos.BaseURL{BucketURL: u}
	Cos = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  core.SecretID,
			SecretKey: core.SecretKey,
		},
	})
}
