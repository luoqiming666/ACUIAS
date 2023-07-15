package main

import (
	"github.com/gin-gonic/gin"
	"test.com/hello/app/controllers"
)

func main() {
	router := gin.Default()

	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	router.Run(":8080")
}
