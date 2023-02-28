package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"summer/config"
	"summer/constant"
	"time"
)

func init() {
	initMysqlConfigure()
}

// initMysqlConfigure 初始化Mysql
func initMysqlConfigure() {
	// 获取Mysql相关配置属性
	dbConfig := &constant.Config.Mysql
	// Mysql配置
	mysqlConfig := getMysqlConfig(dbConfig)
	// gorm 配置
	gormConfig := getGormConfig(dbConfig)
	// 连接数据库
	client, err := gorm.Open(mysql.New(mysqlConfig), gormConfig)
	if err != nil {
		panic(fmt.Errorf("mysql open error: %s", err))
	}
	db, _ := client.DB()
	db.SetMaxIdleConns(10)               //连接池最大空闲数
	db.SetMaxOpenConns(100)              //连接池数据库的最大打开连接数,默认为0不上限.
	db.SetConnMaxLifetime(time.Hour * 4) //设置可以重用连接的最长时间
	// 放到全局变量
	constant.Db = db

}

// 自定义Mysql配置
func getMysqlConfig(conf *config.Mysql) mysql.Config {
	return mysql.Config{
		DSN:                      conf.Dns(), //数据库链接
		DefaultStringSize:        256,        // string 类型字段的默认长度
		DisableDatetimePrecision: true,       //string 类型字段的默认长度
	}
}

// getGormConfig 自定义Gorm配置
func getGormConfig(conf *config.Mysql) *gorm.Config {
	logLevel := logger.Silent
	if conf.LogLevel == "dev" {
		logLevel = logger.Info
	}
	//自定义日志器
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 5, // 慢SQL 阈值
			LogLevel:                  logLevel,        // 日志等级
			IgnoreRecordNotFoundError: true,            //忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,           // 禁用彩色打印
		},
	)
	//gorm 配置
	config := gorm.Config{
		//数据表命名策略  禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		//日志
		Logger: newLogger,
		//禁止自动创建外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
	}
	return &config
}
