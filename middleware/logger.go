/**
  @author: Zero
  @date: 2023/3/18 16:28:40
  @desc: gin整合logrus日志中间件

**/

package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"summer/constant"
	"time"
)

// Logger 自定义Gin日志插件。使用logrus来做
func Logger() gin.HandlerFunc {
	// 新建实例
	logger := logrus.New()
	//设置日志级别
	logger.SetLevel(logrus.DebugLevel)
	//日志分割配置
	logWriter, err := newRotateLogs()
	if err != nil {
		panic(err)
	}
	// 所有日志级别都写到这一个文件即可
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	// 设置日志格式和时间格式 钩子处理器
	hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 使用钩子
	logger.AddHook(hook)
	// 处理中间件
	return func(context *gin.Context) {
		// 开始处理时间
		startTime := time.Now()
		// 处理请求
		context.Next()
		// 请求执行结束时间
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		// 获取主机名
		hostname, err := os.Hostname() //请求ip
		if err != nil {
			hostname = "unknown"
		}
		statusCode := context.Writer.Status() //获取响应状态码
		clientIp := context.ClientIP()        //获取客户端IP
		dataSize := context.Writer.Size()     //携带数据大小
		method := context.Request.Method      //请求类型
		path := context.Request.RequestURI    //请求path
		// 输出日志
		entry := logger.WithFields(logrus.Fields{
			"HostName": hostname,
			"Status":   statusCode,
			"SpendMs":  spendTime,
			"Ip":       clientIp,
			"Method":   method,
			"Path":     path,
			"DataSize": dataSize,
		})
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

// 日志文件切割规则
func newRotateLogs() (*rotatelogs.RotateLogs, error) {
	logConf := constant.Config.Logger
	var filePath = logConf.FilePath // 日志文件存放目录
	//获取时间为日的单位,按照每日进行分割。
	day := time.Duration(logConf.FileSplitDay*24) * time.Hour
	// 获取日志文件最大保存的时间
	maxAge := day * time.Duration(logConf.FileMaxAge)
	write, err := rotatelogs.New(
		filePath+"http/http.%Y-%m-%d.log",                 //gin的请求日志文件名格式
		rotatelogs.WithLinkName(filePath+"http/http.log"), //生成一个当天日志文件的软连接
		rotatelogs.WithRotationTime(day),                  //日志切割时间间隔
		rotatelogs.WithMaxAge(maxAge),                     //日志最长保留时间
	)
	if err != nil {
		return nil, err
	}
	return write, nil
}
