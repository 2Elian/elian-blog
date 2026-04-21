package middleware

import (
	"elian-blog/pkg/response"

	"github.com/gin-gonic/gin"
)

var adminRoles = map[string]bool{
	"admin":  true,
	"editor": true,
}

func RBACAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			response.Forbidden(c, "无权限访问")
			c.Abort()
			return
		}

		roleStr, ok := role.(string)
		if !ok || !adminRoles[roleStr] {
			response.Forbidden(c, "无权限访问")
			c.Abort()
			return
		}

		c.Next()
	}
}