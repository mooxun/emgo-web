package passport

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/requests"
	"github.com/mooxun/emgo-web/pkg/response"
	"github.com/mooxun/emgo-web/routes/middleware/jwt"
	"golang.org/x/crypto/bcrypt"
)

type LoginData struct {
	Username  string    `json:"username" validate:"required" label:"用户名"`
	Password  string    `json:"password" validate:"required" label:"用户密码"`
}

type LoginResult struct {
	AccessToken string `json:"access_token"`
}

func Login(c *gin.Context) {
	var data LoginData
	if ok := requests.Params(c, &data); !ok {
		return
	}

	if ok := requests.Check(c, data); !ok {
		return
	}

	admin, err := models.FindAdmin(data.Username)
	if err != nil || admin == nil {
		response.Err(c, response.ErrMsg{
			Code:  400,
			Error: err,
		})
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(data.Password)); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  402,
			Error: errors.New("密码错误"),
		})
		return
	}

	// 生成Token
	claims := jwt.CustomClaims{
		Id:       admin.Id,
		Username: admin.Username,
		Guard:    "platform",
	}
	if token, err := jwt.GenerateToken(claims); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  401,
			Error: err,
		})
		return
	} else {
		Data := LoginResult{AccessToken: token}
		response.Ok(c, Data)
	}
}