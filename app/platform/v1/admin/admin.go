package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/repository"
	"github.com/mooxun/emgo-web/pkg/requests"
	"github.com/mooxun/emgo-web/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

// @Summary 管理员列表
// @Produce	json
// @Param page int 页码
// @Param page_size int 每页数量
// @Param order string 排序 "id:asc,created_at:desc"
// @Param type string 类型
// @Param start_time time 开始时间
// @Param end_time time 结束时间
// @Router /platform/v1/admin/lists [get]
func Lists(c *gin.Context)  {
	var admins []models.PlatformAdmin
	repository.GetLists(c, &admins, models.PlatformAdmin{})
}

func Detail(c *gin.Context)  {
	repository.GetDetail(c, &models.PlatformAdmin{})
}

func Create(c *gin.Context)  {
	var data models.PlatformAdmin

	if ok := requests.Params(c, &data); !ok {
		return
	}

	if ok := requests.Check(c, data); !ok {
		return
	}

	passwordbyte, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		response.Err(c, response.ErrMsg{
			Code:  600,
			Error: err,
		})
		return
	}
	data.Password = string(passwordbyte)

	if err := models.Create(&data); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  702,
			Error: err,
		})
		return
	} else {
		response.Ok(c, &data)
	}
}

func Update(c *gin.Context)  {

}

func Delete(c *gin.Context)  {
	repository.DeleteItem(c, &models.PlatformAdmin{})
}