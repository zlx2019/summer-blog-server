/**
  @author: Zero
  @date: 2023/3/8 13:09:06
  @desc: 日志相关配置属性实体

**/

package config

// Logger 		日志配置
//
// Level  		日志等级
// Prefix		日志前缀
// FilePath		日志生成的目录名
// ShowLine		是否显示行号
// LogInConsole	是否显示操作文件路径
type Logger struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	FilePath     string `yaml:"file_path"`
	ShowLine     bool   `yaml:"show_line"`
	LogInConsole bool   `yaml:"log_in_console"`
}
