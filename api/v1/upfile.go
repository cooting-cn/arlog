package v1

import (
	"arlog/core"
	log "arlog/core/logs"
	"arlog/core/object"
	"arlog/utils/res"
	"context"
	"github.com/gin-gonic/gin"
)

func UpFile(c *gin.Context) {
	//获取图片数据流
	file, err := c.FormFile("upload")
	if err != nil {
		res.Ask(c, 201, err.Error())
		return
	}
	//设置保存在存储桶中的文件地址，命名为图片名
	name := core.Path + "/" + file.Filename

	//打开文件，并向腾讯云COS上传图片
	f, _ := file.Open()
	//上传字节流
	_, err = object.Cos.Object.Put(context.Background(), name, f, nil)

	if err != nil {
		res.Ask(c, 201, err.Error())
		log.Info("上传对象存储失败")
		return
	}
	res.Ask(c, 200, nil)
	//获取文件链接
	url := object.Cos.Object.GetObjectURL(name).String()
	log.Info("图片上传成功!", url)
}
