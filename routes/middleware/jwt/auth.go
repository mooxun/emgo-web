package jwt

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/response"
)

var jwtSecret = []byte("")

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

type CustomClaims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Guard    string `json:"guard"`
	jwt.StandardClaims
}

// Auth 中间件，检查token
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.Err(c, response.ErrMsg{
				Code:  401,
				Error: errors.New("Access Token is empty"),
			})
			c.Abort()
			return
		}
		token = token[7:]
		// parseToken 解析token包含的信息
		claims, err := ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.Err(c, response.ErrMsg{
					Code:  401,
					Error: err,
				})
			}
			response.Err(c, response.ErrMsg{
				Code:  401,
				Error: err,
			})
			c.Abort()
			return
		}
		if ok := models.CheckAdmin(claims.Id, claims.Username); !ok {
			response.Err(c, response.ErrMsg{
				Code:  401,
				Error: errors.New("admin account error"),
			})
			c.Abort()
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
	}
}

// 生成token
func GenerateToken(claims CustomClaims) (string, error) {
	// token超时时间
	claims.ExpiresAt = time.Now().Add(30 * 24 * time.Hour).Unix()
	claims.Issuer = "emgo-web"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析token
func ParseToken(token string) (*CustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := tokenClaims.Claims.(*CustomClaims); ok && tokenClaims.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}
