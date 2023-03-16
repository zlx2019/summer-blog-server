/**
  @author: Zero
  @date: 2023/3/8 13:09:46
  @desc: 服务配置属性

**/

package properties

// Config 配置文件
// 将settings.yml文件中的配置解析到此结构
type Config struct {
	Server   Server     `yaml:"server"`
	Mysql    DataSource `yaml:"mysql"`
	Pgsql    DataSource `yaml:"pgsql"`
	Logger   Logger     `yaml:"logger"`
	SiteInfo SiteInfo   `yaml:"site_info"`
	QQ       QQ         `yaml:"qq"`
	Email    Email      `yaml:"email"`
	Jwt      Jwt        `yaml:"jwt"`
	QiNiuOss Oss        `yaml:"qi_niu"`
}

// Server 	服务相关配置
type Server struct {
	Host   string `yaml:"host"`    //服务运行的IP
	Port   int    `yaml:"port"`    //服务运行的端口
	Env    string `yaml:"env"`     // 运行环境
	DbType string `yaml:"db_type"` //数据库类型
}
