package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"test.com/hello/app/models"
	"test.com/hello/config"
)

// 公共数据库连接对象
var DB *gorm.DB

func InitDB() {

	// 获取数据源名称
	dsn, err := config.LoadConfig()
	if err != nil {
		print("failed to read config")
		return
	}

	var errConn error

	DB, errConn = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if errConn != nil {
		print("connecting to db err!")
		return
	}

	//迁移数据表
	DB.AutoMigrate(&models.User{})

	print("\n\nsucceed to init DB\n\n")

}
