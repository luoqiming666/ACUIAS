package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// 配置结构体
type Config struct {
	ServerPort string
	ServerHost string
	Username   string
	Password   string
	DbName     string
}

// 加载配置信息
func LoadConfig() (string, error) {

	// 读取配置信息
	viper.SetConfigFile("config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("faild to read config.yaml!", err.Error())
		return "", err
	}

	// 用map方式读取
	dbConfig := viper.GetStringMapString("database")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig["username"],
		dbConfig["password"],
		dbConfig["host"],
		dbConfig["port"],
		dbConfig["db_name"],
	)

	return dsn, nil
}
