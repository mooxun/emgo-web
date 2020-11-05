package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mooxun/emgo-web/routes/index/v1"
)

type Routers func(engine *gin.Engine)

var routers []Routers

func loader(items ...Routers) {
	routers = append(routers, items...)
}

func register() *gin.Engine {
	r := gin.New()

	for _, item := range routers {
		item(r)
	}

	return r
}

func Init() *gin.Engine {
	loader(v1.Router)

	return register()
}
