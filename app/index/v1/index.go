package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/pkg/response"
)

func Index(c *gin.Context) {
	response.Ok(c, nil)
}
