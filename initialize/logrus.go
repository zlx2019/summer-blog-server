/**
  *@author: Zero
  *@date: 2023/3/2 14:16:34
  *@desc: 日志组件初始化
**/

package initialize

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"summer/constant"
)

// 日志颜色
// RED 红色
// GREEN 绿色
// YELLOW 黄色
// BLUE 蓝色
// PURPLE 紫红
// CYAN 青蓝色
// WHITE 白色
const (
	RED    = 31
	GREEN  = 32
	YELLOW = 33
	BLUE   = 34
	PURPLE = 35
	CYAN   = 36
	WHITE  = 37
)

// InitLogger 初始化Log组件
func InitLogger() {
	constant.Log = InitLogrus()
}

// InitLogrus 初始化Logrus实例
func InitLogrus() *logrus.Logger {
	logger := logrus.New()                                  //创建logrus一个实例
	logger.SetOutput(os.Stdout)                             //输出类型
	logger.SetReportCaller(constant.Config.Logger.ShowLine) //开启日志返回函数和行号
	logger.SetFormatter(&LogFormatter{})                    // 日志格式使用自定义的
	// 转换日志级别 string 转 int32
	level, err := logrus.ParseLevel(constant.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level) //日志级别
	InitGlobalLogrus()
	return logger
}

// InitGlobalLogrus 初始化全局Logrus
func InitGlobalLogrus() {
	logrus.SetOutput(os.Stdout)
	logrus.SetReportCaller(constant.Config.Logger.ShowLine)
	logrus.SetFormatter(&LogFormatter{})
	// 转换日志级别
	level, err := logrus.ParseLevel(constant.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

// LogFormatter 重写logrus组件格式化
type LogFormatter struct{}

// Format 日志自定义格式化
func (*LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 根据日志类型,展示选用不同的颜色
	var logLevelColor int
	switch entry.Level {
	case logrus.DebugLevel, logrus.TraceLevel:
		// 调试日志
		logLevelColor = BLUE
	case logrus.WarnLevel:
		// 警告日志
		logLevelColor = YELLOW
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		// 错误日志
		logLevelColor = RED
	default:
		// 正常日志
		logLevelColor = GREEN
	}
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 日志日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	logPrefix := constant.Config.Logger.Prefix
	// 控制台日志输出
	if entry.HasCaller() {
		// 日志输出的所在方法路径
		funcVal := entry.Caller.Function
		// 日志输出的所在文件的相对路径
		basePath := path.Base(entry.Caller.File)
		// 日志输出的代码行号
		line := entry.Caller.Line
		// 文件名:代码行号。在编辑器上生成链接,一键跳转到日志位置.
		fileLine := fmt.Sprintf("%s:%d", basePath, line)
		// 格式:
		// [2023-03-02 15:52:33] [info] [summer/initialize.initMysqlConfigure] gorm.go:43 Gorm Init Success.
		fmt.Fprintf(b, "%s [%s] [\x1b[%dm%s\x1b[0m] [\u001B[%dm%s\u001B[0m] %s %s\n", logPrefix, timestamp, logLevelColor, entry.Level, logLevelColor, funcVal, fileLine, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", logPrefix, timestamp, logLevelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
