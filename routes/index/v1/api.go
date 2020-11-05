package v1

import (
	"github.com/gin-gonic/gin"
	indexV1 "github.com/mooxun/emgo-web/app/index/v1"
)

func Router(e *gin.Engine)  {
	v1 := e.Group("/api/v1")
	{
		v1.GET("index", indexV1.Index)
	}
}
