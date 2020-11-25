package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/app/platform/v1/admin"
	"github.com/mooxun/emgo-web/app/platform/v1/passport"
	"github.com/mooxun/emgo-web/routes/middleware/jwt"
)

func Router(e *gin.Engine) {
	v1 := e.Group("/platform/v1")
	{
		v1.POST("passport/login", passport.Login) // 登录
	}

	v1Auth := e.Group("/platform/v1")
	v1Auth.Use(jwt.Auth())
	{
		// 管理员管理
		v1Auth.GET("admin/lists", admin.Lists)    // 管理员列表
		v1Auth.GET("admin/detail/:id", admin.Detail)  // 管理员详情
		v1Auth.POST("admin/create", admin.Create) // 添加管理员
		v1Auth.POST("admin/update/:id", admin.Update) // 更新管理员
		v1Auth.POST("admin/delete/:id", admin.Delete) // 删除管理员
	}
}
