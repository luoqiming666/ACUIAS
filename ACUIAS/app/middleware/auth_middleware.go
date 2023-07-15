package middleware

import (
	"github.com/gin-gonic/gin"
	"test.com/hello/app/services"
)

func AuthMiddleware(authService *services.AuthService) gin.HandlerFunc {

	return func(c *gin.Context) {
		// 鉴别用户身份逻辑
		// ...

		c.Next()
	}
}
