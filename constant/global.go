package constant

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"summer/config"
)

// 全局变量文件

var (
	Config *config.Config //全局配置属性
	Db     *gorm.DB       //全局数据库连接
	Log    *logrus.Logger //全局日志工具
)
