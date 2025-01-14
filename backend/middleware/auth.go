package middleware

import (
	"backend/database"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定义需要身份验证的路径列表
var authRequiredPaths = []string{
	"/v1/change-password",
	// "/api/v1/moment",
	// 其他需要身份验证的路径
}

// AuthMiddleware 是一个用于身份验证的中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从 Cookie 中获取会话标识符
		sessionID, err := c.Cookie("sessionID")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "获取sessionID失败", "error": err.Error()})
			c.Abort()
			return
		}

		// 从 Redis 中获取用户 ID
		userID, err := database.RedisClient.Get(context.Background(), "session:"+sessionID).Result()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "会话已过期或无效"})
			c.Abort()
			return
		}

		// 将用户名存储在上下文中，方便后续处理函数使用
		c.Set("user_id", userID)
		c.Next()
	}
}
