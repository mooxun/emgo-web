package passport

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/requests"
	"github.com/mooxun/emgo-web/pkg/response"
)

// 管理员注册
func Register(c *gin.Context) {
	data := &models.PlatformAdmin{}

	if err := c.ShouldBindJSON(data); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  707,
			Error: err,
		})
		return
	}

	if ok, errMsg := requests.ValidateRequest(data); !ok {
		response.Err(c, response.ErrMsg{
			Code:   414,
			Error:  errors.New("表单验证失败"),
			Result: errMsg,
		})
		return
	}

	response.Ok(c, data)
}
