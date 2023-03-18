/**
  @author: Zero
  @date: 2023/3/18 11:09:03
  @desc: 自定义输出到控制台的日志格式

**/

package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"summer/constant"
)

// ConsoleLogFormatter 自定义日志输出格式
// 重写logrus的 Format()方法
type ConsoleLogFormatter struct{}

// Format 日志自定义格式化
func (*ConsoleLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
	// 拼接日志
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	// 日志的时间日期格式
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	// 日志的前缀
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
		// [2023-03-02 15:52:33] [info] [summer/initialize.initDataSourceConfigure] gorm.go:43 Gorm Init Success.
		fmt.Fprintf(b, "%s [%s] [\x1b[%dm%s\x1b[0m] [\u001B[%dm%s\u001B[0m] %s %s\n", logPrefix, timestamp, logLevelColor, entry.Level, logLevelColor, funcVal, fileLine, entry.Message)
	} else {
		fmt.Fprintf(b, "%s [%s] \x1b[%dm[%s]\x1b[0m %s\n", logPrefix, timestamp, logLevelColor, entry.Level, entry.Message)
	}
	return b.Bytes(), nil
}
