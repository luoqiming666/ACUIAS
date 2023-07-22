package main

import (
	"github.com/gin-gonic/gin"
	"test.com/hello/app/controllers"
	"test.com/hello/app/middleware"
	"test.com/hello/database"
)

func main() {

	// 连接数据库
	database.InitDB()

	router := gin.Default()

	// 管理者
	c := &controllers.ControllerManager{}
	m := &middleware.MiddlewareManager{}

	// 路由中间件
	apiGroup := router.Group("/api")
	// apiGroup.POST("/login", middleware.ValidateTokenMiddleware(), c.Login)
	apiGroup.POST("/login", c.Login)
	apiGroup.POST("/register", c.Register)

	// 需要中间件
	apiGroup.GET("/userinfo", m.ValidateTokenMiddleware(), c.GetUserInfo)

	router.Run(":8080")
}
