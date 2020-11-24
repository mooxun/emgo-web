package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/app/platform/v1/passport"
)

func Router(e *gin.Engine) {
	v1 := e.Group("/platform/v1")
	{
		v1.POST("passport/register", passport.Register)	// 管理员注册
	}
}
