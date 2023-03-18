/**
  *@author: Zero
  *@date: 2023/3/2 14:16:34
  *@desc: 日志组件初始化
**/

package initialize

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"summer/constant"
	"summer/handler/log"
	"time"
)

// InitLogger 初始化Log组件
func InitLogger() {
	constant.Log = InitLogrus()
}

// InitLogrus 初始化Logrus实例
func InitLogrus() *logrus.Logger {
	//创建logrus一个实例
	logger := logrus.New()
	// 创建日志切割配置
	logFileWrite, err := newRotateLogs()
	if err != nil {
		panic(err)
	}
	//开启日志返回函数和行号
	logger.SetReportCaller(constant.Config.Logger.ShowLine)
	//TODO 因为终端日志部分内容 携带了color,所以要和日志文件采用不同的日志格式

	// 设置自定义的终端日志格式
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.ConsoleLogFormatter{})
	// 设置要输出的日志文件，和日志文件输出格式
	logger.AddHook(NewWriterHook(logFileWrite, &log.FileLogFormatter{}))
	// 根据配置的项目日志级别,开启logrus的日志级别
	level, err := logrus.ParseLevel(constant.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logger.SetLevel(level) //日志级别
	InitGlobalLogrus(logFileWrite)
	return logger
}

// InitGlobalLogrus 初始化全局Logrus
func InitGlobalLogrus(logFileWrite *rotatelogs.RotateLogs) {
	logrus.SetReportCaller(constant.Config.Logger.ShowLine)
	// 设置自定义的终端日志格式
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&log.ConsoleLogFormatter{})
	// 设置日志文件输出格式
	logrus.AddHook(NewWriterHook(logFileWrite, &log.FileLogFormatter{}))
	// 转换日志级别
	level, err := logrus.ParseLevel(constant.Config.Logger.Level)
	if err != nil {
		level = logrus.InfoLevel
	}
	logrus.SetLevel(level)
}

// WriterHook 自定义Hook
// 用于将日志输入到不同Writer中时使用不同的日志格式
type WriterHook struct {
	Writer    io.Writer
	Formatter logrus.Formatter
}

// NewWriterHook 构建WriterHook
// writer: 要写入的日志文件
// formatter 日志文件的格式
func NewWriterHook(writer io.Writer, formatter logrus.Formatter) *WriterHook {
	return &WriterHook{
		Writer:    writer,
		Formatter: formatter,
	}
}

// Levels 设置哪些级别的日志会触发此钩子函数
func (hook *WriterHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire 日志数据写入到指定writer
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := hook.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write(line)
	return err
}

// 日志文件切割规则
func newRotateLogs() (*rotatelogs.RotateLogs, error) {
	logConf := constant.Config.Logger
	// 日志文件存放目录
	var filePath = logConf.FilePath
	//获取时间为日的单位,按照每日进行分割。
	day := time.Duration(logConf.FileSplitDay*24) * time.Hour
	// 获取日志文件最大保存的时间
	maxAge := day * time.Duration(logConf.FileMaxAge)
	write, err := rotatelogs.New(
		filePath+"business/business.%Y-%m-%d.log",
		rotatelogs.WithLinkName(filePath+"business/business.log"), //生成一个当天日志文件的软连接
		rotatelogs.WithRotationTime(day),                          //日志切割时间间隔
		rotatelogs.WithMaxAge(maxAge),                             //日志最长保留时间
	)
	if err != nil {
		return nil, err
	}
	return write, nil
}
