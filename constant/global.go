package constant

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"summer/config"
)

// 全局变量文件

var (
	Config *config.Config //全局配置属性
	Db     *sql.DB        //全局数据库连接
	Log    *logrus.Logger //全局日志工具
)
