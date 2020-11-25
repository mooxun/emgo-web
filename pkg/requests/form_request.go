package requests

import (
	"errors"
	"log"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/mooxun/emgo-web/pkg/response"
)

// 表单验证
// https://github.com/go-playground/validator
var (
	uni           *ut.UniversalTranslator
	Validate      *validator.Validate
	ValidateTrans ut.Translator
)

func init() {
	// validator 中文错误消息
	zhCn := zh.New()
	uni = ut.New(zhCn, zhCn)

	// 验证器注册中文错误消息
	ValidateTrans, _ = uni.GetTranslator("zh")

	// 验证器
	Validate = validator.New()
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("title")
	})
	if err := zh_translations.RegisterDefaultTranslations(Validate, ValidateTrans); err != nil {
		log.Println(err.Error())
	}
}

// 验证表单数据
func ValidateRequest(data interface{}) (ok bool, errMsg []string) {
	err := Validate.Struct(data)
	if err != nil {
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errMsg = append(errMsg, e.Translate(ValidateTrans))
		}
	} else {
		ok = true
	}
	return
}

// 响应表单验证结果
func Check(c *gin.Context, data interface{}) bool {
	if ok, errMsg := ValidateRequest(data); !ok {
		response.Err(c, response.ErrMsg{
			Code:   414,
			Error:  errors.New("表单验证失败"),
			Result: errMsg,
		})
		return false
	}
	return true
}
