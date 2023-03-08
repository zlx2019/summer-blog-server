/**
  @author: Zero
  @date: 2023/3/8 13:09:46
  @desc: Mysql配置属性实体

**/

package config

import "fmt"

// DataSource 配置
// LogLevel Sql日志等级. debug:输出全部sql。dev: 开发环境。release: 生产环境。
type DataSource struct {
	Host        string `yaml:"host"`
	Port        int    `yaml:"port"`
	DbName      string `yaml:"db_name"`
	Username    string `yaml:"username"`
	Password    string `yaml:"password"`
	LogLevel    string `yaml:"log_level"`
	CreateTable bool   `yaml:"create_table"`
}

// MysqlDns 拼接Mysql连接
func (m *DataSource) MysqlDns() string {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.DbName)
}

// PgSqlDns 拼接PgSql连接
func (d *DataSource) PgSqlDns() string {
	// host=ip port=端口 user=数据库角色 password=角色密码 dbname=数据库名称 sslmode=disable TimeZone=Asia/Shanghai
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		d.Host, d.Port, d.Username, d.Password, d.DbName)
}
