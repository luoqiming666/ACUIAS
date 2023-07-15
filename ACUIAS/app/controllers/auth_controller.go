package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/hello/app/services"
)

// 登录
func Login(c *gin.Context) {

	// 创建一个区块链服务
	AS := services.BlockchainService{}

	//区块链认证服务
	authService := services.NewAuthService(AS)

	// 调用身份认证服务的登录方法进行用户身份验证
	// ...
	// 判断服务处理结果

	print(authService)

	c.JSON(http.StatusOK, gin.H{"message": "pass，Login successful"})
}

// 注册
func Register(c *gin.Context) {

	// 创建一个区块链服务
	AS := services.BlockchainService{}

	//区块链认证服务
	authService := services.NewAuthService(AS)

	// 调用身份认证服务的注册方法进行用户身份验证
	// ...

	print(authService)

	c.JSON(http.StatusCreated, gin.H{"message": "pass，Registration successful"})
}
