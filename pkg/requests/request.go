package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/pkg/response"
)

// 绑定请求参数
func Params(c *gin.Context, params interface{}) bool {
	if err := c.ShouldBindJSON(&params); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  707,
			Error: err,
		})
		return false
	}
	return true
}
