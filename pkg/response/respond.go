package response

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/ecode"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

type ErrMsg struct {
	Code   int
	Error  error
	Result interface{}
}

// 分页列表结果
type ListsResult struct {
	Field      interface{} `json:"field"`
	Data       interface{} `json:"data"`
	Pagination interface{} `json:"pagination"`
}

// 成功响应
func Ok(c *gin.Context, data interface{}) {
	c.JSON(200, Response{
		Code:    0,
		Message: "success",
		Result:  data,
	})
}

// 错误响应
func Err(c *gin.Context, errMsg ErrMsg) {
	var message string
	if !cfg.App.Debug || errMsg.Error == nil {
		message = ecode.GetMsg(errMsg.Code)
	} else {
		message = errMsg.Error.Error()
	}

	c.JSON(200, Response{
		Code:    errMsg.Code,
		Message: message,
		Result:  errMsg.Result,
	})
}
