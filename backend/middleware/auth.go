package middleware

import (
	"backend/database"
	"context"

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
			// 即使获取sessionID失败，也不直接返回错误，而是继续后续处理
			c.Set("user_id", "")
			c.Next()
			return
		}

		// 从 Redis 中获取用户 ID
		userID, err := database.RedisClient.Get(context.Background(), "session:"+sessionID).Result()
		if err != nil {
			// 即使会话已过期或无效，也不直接返回错误，而是继续后续处理
			c.Set("user_id", "")
			c.Next()
			return
		}

		// 将用户名存储在上下文中，方便后续处理函数使用
		c.Set("user_id", userID)
		c.Next()
	}
}
