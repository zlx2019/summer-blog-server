package constant

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"summer/properties"
)

// 全局变量文件

var (
	Config *properties.Config //全局配置属性
	Db     *gorm.DB           //全局数据库连接
	Log    *logrus.Logger     //全局日志工具
)

const (
	// ConfigFile 配置文件路径
	ConfigFile = "settings.yml"
)
