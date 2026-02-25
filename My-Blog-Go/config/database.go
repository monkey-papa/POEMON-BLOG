// Package config 数据库配置
package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB 全局数据库连接
var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
	)

	// 配置日志级别
	// DB_LOG_SQL=true 时显示所有SQL日志，否则只显示错误
	logLevel := logger.Silent
	if AppConfig.DBLogSQL {
		logLevel = logger.Info
	} else if AppConfig.Environment != "production" {
		logLevel = logger.Error // 开发环境默认只显示错误
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		// 禁用默认事务（提升性能）
		SkipDefaultTransaction: true,
		// 预编译语句
		PrepareStmt: true,
	})
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 获取底层sql.DB对象进行连接池配置
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("获取数据库连接失败:", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大存活时间

	log.Println("数据库连接成功")
}
