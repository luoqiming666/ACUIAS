package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"test.com/hello/app/models"
	"test.com/hello/app/services"
	"test.com/hello/utils"
)

// 控制器管理
type ControllerManager struct {
	userServ services.Userservice
}

// 登录
func (cm *ControllerManager) Login(c *gin.Context) {

	var userPost models.User

	ut := &utils.Userutils{}

	if err := c.ShouldBind(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errorPost": err.Error(),
		})
		return
	}

	// 判断用户名是否为空
	if userPost.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "username not allowed to be null",
		})
		return
	}

	// 判断密码是否为空
	if userPost.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "password not allowed to be null",
		})
		return
	}

	// 查询用户
	userGet, _ := cm.userServ.GetUserService(userPost.Username)
	if userGet.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "username not correct",
		})
		return

	} else {

		// 登录成功
		if ut.VerifyPassword(userPost.Password, userGet.Password) {

			c.JSON(http.StatusOK, gin.H{
				"msg": "password correct, login success!",
			})

			// 服务器生成并存储随机token
			if randomToken, isGenerate := cm.userServ.UpdateToken(userGet); isGenerate == true {
				c.JSON(http.StatusOK, gin.H{
					"msg": "success to generate random token:" + randomToken,
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"msg": "failed to generate random token",
				})

			}

		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": "password not correct",
			})

		}

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

	ut := &utils.Userutils{}

	if err := c.ShouldBind(userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "binding post error!",
		})
		return
	}

	//判断合法性

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

	//名字是否为空
	if userPost.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Username not allow to be null",
		})
		return
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

	// 加密密码
	encodedPassword := ut.EncryptWithMD5(userPost.Password)

	userPost.Password = encodedPassword

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

		//

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

// 获取用户个人信息
func (cm *ControllerManager) GetUserInfo(c *gin.Context) {

	// 判断权限
	if userIDRow, IDexist := c.Get("userID"); !IDexist {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "key doesn't existed",
		})
		return
	} else {

		if userID, ok := userIDRow.(float64); !ok {

			fmt.Println("failed to get userID")
			return

		} else {
			// 获取用户信息
			if getUser, err := cm.userServ.GetAllUserInfoByID(int(userID)); err != nil {
				print(getUser)
			} else {
				fmt.Println(getUser.Username, getUser.Token)
			}

		}

	}

}
