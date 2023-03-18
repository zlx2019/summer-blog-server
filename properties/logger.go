/**
  @author: Zero
  @date: 2023/3/8 13:09:06
  @desc: 日志相关配置属性实体

**/

package properties

// Logger 		日志配置
//
// Level  		日志等级
// Prefix		日志前缀
// FilePath		日志生成的目录名
// ShowLine		是否显示行号
// LogInConsole	是否显示操作文件路径
type Logger struct {
	Level        string `yaml:"level"`          //日志级别 [panic fatal error warn info debug trace]
	Prefix       string `yaml:"prefix"`         // 日志格式前缀
	ShowLine     bool   `yaml:"show_line"`      // 是否显示代码行号
	LogInConsole bool   `yaml:"log_in_console"` //是否显示操作文件路径
	FilePath     string `yaml:"file_path"`      //生成的日志文件目录
	FileSplitDay int    `yaml:"file_split_day"` //每几天生成一个新的日志文件
	FileMaxAge   int    `yaml:"file_max_age"`   //日志文件最多保留几天
}
