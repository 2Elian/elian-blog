package middleware

import (
	"elian-blog/internal/utils"
	"elian-blog/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			response.Unauthorized(c, "未提供认证token")
			c.Abort()
			return
		}

		if strings.HasPrefix(tokenStr, "Bearer ") {
			tokenStr = tokenStr[7:]
		}

		claims, err := utils.ParseToken(tokenStr, secret)
		if err != nil {
			response.Unauthorized(c, "token无效或已过期")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}
