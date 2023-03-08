package config

// Config 配置文件
// 将settings.yml文件中的配置解析到此结构
type Config struct {
	Server Server     `yaml:"server"`
	Mysql  DataSource `yaml:"mysql"`
	Pgsql  DataSource `yaml:"pgsql"`
	Logger Logger     `yaml:"logger"`
}

// Server 	服务相关配置
//
// Host 	服务运行的IP
// Port 	服务运行的端口
// Env 		运行环境
// DbType	数据库类型
type Server struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	Env    string `yaml:"env"`
	DbType string `yaml:"db_type"`
}
