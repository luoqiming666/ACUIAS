package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"test.com/hello/app/models"
	"test.com/hello/app/services"
)

// 控制器管理
type ControllerManager struct {
	userServ services.Userservice
}

// 登录
func (cm *ControllerManager) Login(c *gin.Context) {

	var user models.User

	// 获取用户表单内容
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorPost": err.Error(),
		})
		return
	}

	if user.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "please input a username",
		})
		return
	}

	// ------------------------区块链业务部分----------------------------

	// // 创建一个区块链服务
	// AS := services.BlockchainService{}

	// //区块链认证服务
	// authService := services.NewAuthService(AS)

	// // 调用身份认证服务的登录方法进行用户身份验证
	// // ...
	// // 判断服务处理结果

	// print(authService)

	// c.JSON(http.StatusOK, gin.H{"message": "pass，Login successful"})
}

// 注册
func (cm *ControllerManager) Register(c *gin.Context) {

	userPost := &models.User{}

	if err := c.ShouldBind(userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "binding post error!",
		})
		return
	}

	// 检测空名

	//检查重名
	if cm.userServ.IsUsernameExist(userPost.Username) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "user already exist",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg": "posting username is unique",
		})
	}

	//密码是否为空
	if userPost.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "password not allow to be null",
		})
		return
	}

	//手机是否为空
	if userPost.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "phone not allow to be null",
		})
		return
	}

	//邮箱是否为空
	if userPost.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "email not allow to be null",
		})
		return
	}

	// 新增用户
	if _, errCreate := cm.userServ.CreateUserService(userPost); errCreate != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error occur when creating users",
		})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{
			"msg": "register success",
		})
		return

	}

	// 区块链业务部分

	// // 创建一个区块链服务
	// AS := services.BlockchainService{}

	// //区块链认证服务
	// authService := services.NewAuthService(AS)

	// // 调用身份认证服务的注册方法进行用户身份验证
	// // ...

	// print(authService)

	// c.JSON(http.StatusCreated, gin.H{"message": "pass，Registration successful"})
}
