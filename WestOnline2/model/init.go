package model

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DataBase(dsn string) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Mysql数据库连接错误")
	} else {
		fmt.Println("mysql数据库连接成功")
	}

	DB = db
	migration()
}
