package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, x-terminal-token, x-terminal-id")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		gin.DefaultWriter.Write([]byte(
			"[" + start.Format("2006-01-02 15:04:05") + "] " +
				method + " " + path + " " +
				"status=" + intToStr(status) + " " +
				"latency=" + latency.String() + "\n",
		))
	}
}

func intToStr(i int) string {
	if i >= 100 {
		return string([]byte{byte(i/100 + '0'), byte(i%100/10 + '0'), byte(i%10 + '0')})
	}
	if i >= 10 {
		return string([]byte{byte(i/10 + '0'), byte(i%10 + '0')})
	}
	return string([]byte{byte(i + '0')})
}
