package logs

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func InitLog() {
	Log = logrus.New()
	// 为当前logrus实例设置消息的输出，同样地，
	// 可以设置logrus实例的输出到任意io.writer
	Log.Out = os.Stdout
	// 为当前logrus实例设置消息输出格式为json格式。

	//设置在输出日志中添加文件名和方法信息
	//Log.SetReportCaller(true)
	//日志配置
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:03:04",

		/*		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
				//处理文件名
				fileName := path.Base(frame.File)
				return frame.Function, fileName
			},*/
	})

}

func Info(v ...any) {
	Log.Info(v)
}

func Error(v ...any) {
	Log.Error(v)
}

func Debug(v ...any) {
	Log.Debug(v)
}
func Panic(v ...any) {
	Log.Panic(v)
}

func LoggerToFile() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqUri := c.Request.RequestURI
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		//日志格式
		Log.WithFields(logrus.Fields{
			"http_status": statusCode,
			"total_time":  latencyTime,
			"ip":          clientIP,
			"method":      reqMethod,
			"uri":         reqUri,
		}).Info("ok")
	}
}
