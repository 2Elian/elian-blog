package middleware

import (
	"elian-blog/internal/dao"
	"elian-blog/internal/model"

	"github.com/gin-gonic/gin"
)

func OperationLog(logDao *dao.OperationLogDao) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		userID, _ := c.Get("user_id")
		uid, _ := userID.(uint)

		log := &model.OperationLog{
			UserID: uid,
			Method: c.Request.Method,
			URL:    c.Request.URL.Path,
			IP:     c.ClientIP(),
			Desc:   c.Request.Method + " " + c.Request.URL.Path,
		}

		go func() {
			_ = logDao.Create(log)
		}()
	}
}
