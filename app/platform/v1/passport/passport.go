package passport

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/response"
	"go.mongodb.org/mongo-driver/bson"
)

func Test(c *gin.Context) {
	var user []models.PlatformAdmin
	_ = models.PlColl().Find(context.Background(), bson.M{}).All(&user)
	response.Ok(c, user)
}

func Register(c *gin.Context) {
	data := &models.PlatformAdmin{}
	if err := c.ShouldBindJSON(data); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  707,
			Error: err,
		})
		return
	}

	res, err := models.PlColl().InsertOne(context.Background(), data)
	if err != nil {
		response.Err(c, response.ErrMsg{
			Code:  707,
			Error: err,
		})
	}

	response.Ok(c, res)
}
