package main

import (
	"github.com/gin-gonic/gin"
	"test.com/hello/app/controllers"
	"test.com/hello/database"
)

func main() {

	// 连接数据库
	database.InitDB()

	router := gin.Default()

	c := &controllers.ControllerManager{}

	router.POST("/login", c.Login)
	router.POST("/register", c.Register)
	// router.POST("/del")

	router.Run(":8080")
}
