package model

import (
	"fmt"
	"ginblog/utils"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDb() {
	a := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassword,
		utils.DbHost,
		utils.DbPort,
		utils.Dbname,
	)
	fmt.Println(a)
	db, err = gorm.Open(mysql.Open(a), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败", err)
	}

	db.AutoMigrate(&User{}, &Category{}, &Article{})

	sqlDB, _ := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Second * 10)

}
