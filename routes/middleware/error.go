package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/pkg/logger"
)

// 捕获到在http处理时的错误
// 在 handler 和其他地方如果产生了 error 可直接panic，到这里统一处理，简化 if err != nil 之类的代码
func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("%v", err)
				c.Abort()
				c.JSON(200, err)
			}
		}()
		c.Next()
	}
}
