package config

import "fmt"

// Config 配置文件
// 将settings.yml文件中的配置解析到此结构
type Config struct {
	Server Server `yaml:"server"`
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
}

// Server 	服务相关配置
//
// Host 	服务运行的IP
// Port 	服务运行的端口
// Env 		运行环境
type Server struct {
	// 服务运行的IP
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

// Mysql 配置
// LogLevel Sql日志等级. debug:输出全部sql。dev: 开发环境。release: 生产环境。
type Mysql struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DbName      string `yaml:"db_name"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	LogLevel    string `yaml:"log_level"`
	CreateTable bool   `yaml:"create_table"`
}

// Dns 拼接Mysql链接
func (m *Mysql) Dns() string {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DbName)
}

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
