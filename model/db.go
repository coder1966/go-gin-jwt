package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"go-gin-jwt/util"
	"time"
)

// 入口文件，描述在数据库的参数的

var db *gorm.DB // 指定原型
var err error   // 接收错误信息

func InitDB() {
	db, err = gorm.Open(util.DB, fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		util.DBUser,
		util.DBPassword,
		util.DBHost,
		util.DBPort,
		util.DBName,
	))
	if err != nil {
		fmt.Printf("连接数据库失败。err=：%s", err)
	}

	// 自动迁移。自动创建数据表，数据库表名自动加复数。// db.SingularTable(true) // 数据库表名不自动加复数
	db.AutoMigrate(
		&User{},
		&Article{},
		&Category{},
	)

	// 连接池
	// SetMaxIdleCons 设置连接池中的最大闲置连接数。
	db.DB().SetMaxIdleConns(10)
	// SetMaxOpenCons 设置数据库的最大连接数量。
	db.DB().SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。数据库连接时间要小于r.Run(utils.HttpPort)框架连接时间
	db.DB().SetConnMaxLifetime(10 * time.Second)

	//db.Close()
}
