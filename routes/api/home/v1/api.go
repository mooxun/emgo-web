package v1

import (
	"github.com/gin-gonic/gin"
	homeV1 "github.com/mooxun/emgo-web/app/home/v1"
)

func Router(e *gin.Engine)  {
	v1 := e.Group("/api/v1")
	{
		v1.GET("index", homeV1.Index)
	}
}
