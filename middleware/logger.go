package middleware

// 自定义的日志中间件
import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"go-gin-jwt/util"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	//LogName := time.Now().Format("2006-01-02")
	//logPath := utils.LogPath + LogName + ".log"                   // 日志文件名。每天一个文件
	logPath := util.LogPath + "log"
	linkName := "latest_Log.log"                                  // 软连接，最新的文件
	src, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE, 0755) // 打开|创建文件，权限0755.
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New() // 实例化日志
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := rotatelogs.New(
		logPath+"%Y%m%d.log",                      // 文件名
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 7*24小时记录
		rotatelogs.WithRotationTime(24*time.Hour), // 24小时分割一次
		rotatelogs.WithLinkName(linkName),         // 软连接，根目录下可以查询最新的日志，最新的文件。windows 用软连接，需要管理员权限
	)

	// 7种日志都写在那个文件里
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}

	// 时间格式化
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook) // 实例化一个hook

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next() // 洋葱模型的中间件，这里跳转去执行下一个中间件。
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0))) // 花费的时间
		hostName, err := os.Hostname()                                                               // 请求主机。客户端请求过来的名字
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()    // 状态码
		clientIp := c.ClientIP()           // 客户端的IP
		userAgent := c.Request.UserAgent() // 客户端的浏览器型号、端口等信息。
		dataSize := c.Writer.Size()        // 请求过来文件的大小
		if dataSize < 0 {
			dataSize = 1
		}
		method := c.Request.Method   // 请求的方法 GET POST ......
		path := c.Request.RequestURI // 请求的路径

		// “github.com/sirupsen/logrus” 作者建议的标准写法
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String()) // 如果系统内部有错误。把错误码记录。
		}
		if statusCode >= 500 {
			entry.Error() // 错误
		} else if statusCode >= 400 {
			entry.Warn() // 警告
		} else {
			entry.Info() // 普通信息。一般是成功信息。
		}

	}
}
