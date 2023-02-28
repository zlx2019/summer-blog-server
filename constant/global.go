package constant

import (
	"database/sql"
	"summer/config"
)

// 全局变量文件

var (
	Config *config.Config //全局配置属性
	Db     *sql.DB        //全局数据库连接
)
