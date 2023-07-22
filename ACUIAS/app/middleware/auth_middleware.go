package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"test.com/hello/app/services"
	"test.com/hello/utils"
)

// 中间件
type MiddlewareManager struct {
}

// 请求头判断token
func (mid *MiddlewareManager) ValidateTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取Token
		tokenString := c.GetHeader("Authorization")

		// 没有token
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
			return
		}

		// 解析并验证Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 指定使用相同的密钥进行签名验证
			return utils.SigningKey, nil
		})

		// 解析的token不合法
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg": "token is valid!",
			})
		}

		// 将解析后的Token信息存储在上下文中，供后续处理函数使用
		claims := token.Claims.(jwt.MapClaims)
		c.Set("userID", claims["userID"])
		c.Next()
	}
}

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {

	return func(c *gin.Context) {
		// 鉴别用户身份逻辑
		// ...

		c.Next()
	}
}
