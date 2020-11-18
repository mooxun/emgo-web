package routes

import (
	"github.com/gin-gonic/gin"
	homeV1 "github.com/mooxun/emgo-web/routes/api/home/v1"
	platformV1 "github.com/mooxun/emgo-web/routes/api/platform/v1"
	"github.com/mooxun/emgo-web/routes/middleware"
)

type Routers func(engine *gin.Engine)

var routers []Routers

// 加载路由
func loader(items ...Routers) {
	routers = append(routers, items...)
}

// 路由注册
func register() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.ErrorHandler(), middleware.Cors())

	for _, item := range routers {
		item(r)
	}

	return r
}

func Init() *gin.Engine {
	loader(homeV1.Router, platformV1.Router)

	return register()
}
