package v1

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	panic("errer")
	c.JSON(200, "app api v1 index")
}
