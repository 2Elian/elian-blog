package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimit(rps int) gin.HandlerFunc {
	limiter := rate.NewLimiter(rate.Every(time.Second/time.Duration(rps)), rps*2)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.AbortWithStatusJSON(429, gin.H{"code": 429, "message": "请求过于频繁"})
			return
		}
		c.Next()
	}
}
