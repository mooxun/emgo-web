package routes

import (
	"net/http"

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
	r.StaticFile("/", "emgo-admin/index.html")
	r.StaticFile("/color.less", "emgo-admin/color.less")
	r.StaticFile("/favicon.ico", "emgo-admin/favicon.ico")
	r.Static("/css", "emgo-admin/css")
	r.Static("/js", "emgo-admin/js")
	r.Static("/img", "emgo-admin/img")
	r.Static("/loading", "emgo-admin/loading")

	r.LoadHTMLFiles("emgo-admin/index.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

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
