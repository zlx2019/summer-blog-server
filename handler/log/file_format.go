/**
  @author: Zero
  @date: 2023/3/18 11:07:36
  @desc: 自定义输出到日志文件的日志格式

**/

package log

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"path"
	"summer/constant"
)

// FileLogFormatter 自定义日志输出格式
// 重写logrus的 Format()方法
type FileLogFormatter struct{}

// Format 日志自定义格式化
func (*FileLogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
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
	// 日志输出的所在方法路径
	funcVal := entry.Caller.Function
	// 日志输出的所在文件的相对路径
	basePath := path.Base(entry.Caller.File)
	// 日志输出的代码行号
	line := entry.Caller.Line
	// 文件名:代码行号。在编辑器上生成链接,一键跳转到日志位置.
	fileLine := fmt.Sprintf("%s:%d", basePath, line)
	// 格式(不包含颜色):
	// [2023-03-02 15:52:33] [info] [summer/initialize.initDataSourceConfigure] gorm.go:43 Gorm Init Success.
	fmt.Fprintf(b, "%s [%s] [%s] [%s] %s %s\n", logPrefix, timestamp, entry.Level, funcVal, fileLine, entry.Message)
	return b.Bytes(), nil
}
