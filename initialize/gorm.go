package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"summer/config"
	"summer/constant"
	. "summer/models"
	"summer/utils"
	"time"
)

// initDataSourceConfigure 初始化Mysql
func initDataSourceConfigure() {
	// 获取Mysql相关配置属性
	dbConfig := &constant.Config.Mysql
	if utils.StrIsBlank(dbConfig.Host) {
		constant.Log.Panic("DataSource Host 未配置")
	}
	var dial gorm.Dialector
	var gormConfig *gorm.Config
	if constant.Config.Server.DbType == "pgsql" {
		// 配置pgsql连接
		dial = postgres.New(getPostgresSqlConfig(&constant.Config.Pgsql))
		gormConfig = getGormConfig(&constant.Config.Pgsql)
	} else {
		// 配置mysql连接
		dial = mysql.New(getMysqlConfig(&constant.Config.Mysql))
		gormConfig = getGormConfig(&constant.Config.Mysql)
	}
	// 连接数据库
	client, err := gorm.Open(dial, gormConfig)
	if err != nil {
		panic(fmt.Errorf("mysql open error: %s", err))
	}
	db, _ := client.DB()
	db.SetMaxIdleConns(10)               //连接池最大空闲数
	db.SetMaxOpenConns(100)              //连接池数据库的最大打开连接数,默认为0不上限.
	db.SetConnMaxLifetime(time.Hour * 4) //设置可以重用连接的最长时间
	// 放到全局变量
	constant.Db = client
	constant.Log.Info("Gorm Init Success.")
	// 表结构迁移生成
	if dbConfig.CreateTable {
		tableMigrate(client)
	}
}

// 根据结构生成表结构
func tableMigrate(client *gorm.DB) {
	// 设置自定义关联表 替代默认生成的关系表
	client.SetupJoinTable(&User{}, "LikeArticles", &UserLikeArticle{})
	client.SetupJoinTable(&Article{}, "tags", &ArticleTag{})
	client.SetupJoinTable(&Menu{}, "Banners", &MenuFile{})

	// 根据结构自动创建数据表
	err := client.AutoMigrate(
		&User{},
		&Article{},
		&UserLikeArticle{},
		&Tag{},
		&ArticleTag{},
		&Comment{},
		&FadeBack{},
		&File{},
		&Menu{},
		&MenuFile{},
		&Message{},
		&Advert{},
		&LoginLog{},
	)
	if err != nil {
		constant.Log.Fatalf("表结构生成失败: %s", err)
	}
	constant.Log.Info("Table Create Success")
}

// 获取自定义Mysql配置
func getMysqlConfig(conf *config.DataSource) mysql.Config {
	return mysql.Config{
		DSN:                      conf.MysqlDns(), //数据库链接
		DefaultStringSize:        256,             // string 类型字段的默认长度
		DisableDatetimePrecision: true,            //string 类型字段的默认长度
	}
}

// 获取自定义PgSql配置
func getPostgresSqlConfig(conf *config.DataSource) postgres.Config {
	return postgres.Config{
		DSN:                  conf.PgSqlDns(),
		PreferSimpleProtocol: false,
	}
}

// getGormConfig 自定义Gorm配置
func getGormConfig(conf *config.DataSource) *gorm.Config {
	logLevel := logger.Info
	if conf.LogLevel == "prod" {
		logLevel = logger.Error
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
