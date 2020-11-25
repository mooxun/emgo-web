package repository

import (
	"errors"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mooxun/emgo-web/models"
	"github.com/mooxun/emgo-web/pkg/orm"
	"github.com/mooxun/emgo-web/pkg/requests"
	"github.com/mooxun/emgo-web/pkg/response"
	"github.com/mooxun/emgo-web/pkg/util"
)

// 前端antv表头字段
type ListFields struct {
	DataIndex string `json:"dataIndex"`
	Title     string `json:"title"`
}

// 获取前端antv表头字段
func GetListFields(m interface{}) (f []ListFields) {
	s := reflect.TypeOf(m)
	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Tag.Get("title") != "" {
			f = append(f, ListFields{
				DataIndex: s.Field(i).Tag.Get("json"),
				Title:     s.Field(i).Tag.Get("title"),
			})
		}
	}
	return
}

func GetSelectFields(m interface{}, t string) (f []string) {
	s := reflect.TypeOf(m)
	for i := 0; i < s.NumField(); i++ {
		if s.Field(i).Tag.Get("show") != "" {
			showType := util.Explode(",", s.Field(i).Tag.Get("show"))
			if ok := util.InArray(t, showType); ok {
				f = append(f, s.Field(i).Tag.Get("json"))
			}
		}
	}
	return
}

// 获取全部记录
func GetAll(c *gin.Context, model interface{}) {
	if err := orm.DB.Find(model).Error; err != nil {
		response.Err(c, response.ErrMsg{
			Code:  700,
			Error: err,
		})
		return
	}
	response.Ok(c, model)
}

// @Summary 获取列表
// @Produce	json
// @Param page int 页码
// @Param page_size int 每页数量
// @Param order string 排序 "id:asc,created_at:desc"
// @Param type string 类型
// @Param start_time time 开始时间
// @Param end_time time 结束时间
func GetLists(c *gin.Context, lists interface{}, model interface{}) {
	query := models.Query{
		Ctx:         c,
		DB:          orm.DB,
		SelectField: GetSelectFields(model, "lists"),
	}

	err := query.Fetch(lists)

	if err != nil {
		response.Err(c, response.ErrMsg{
			Code:  700,
			Error: err,
		})
		return
	}

	field := GetListFields(model)
	response.Ok(c, response.ListsResult{
		Field:      field,
		Data:       lists,
		Pagination: query.Pagination,
	})
}

// @Summary 获取详情
func GetDetail(c *gin.Context, model interface{}) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id <= 0 {
		response.Err(c, response.ErrMsg{Code: 707})
	}

	if err := models.Find(id, model); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  700,
			Error: err,
		})
	} else {
		response.Ok(c, &model)
	}
}

// @Summary 创建一条记录
func CreateItem(c *gin.Context, model interface{}) {

	if ok := requests.Params(c, &model); !ok {
		return
	}

	if ok := requests.Check(c, model); !ok {
		return
	}

	if err := models.Create(model); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  702,
			Error: err,
		})
		return
	} else {
		response.Ok(c, &model)
	}
}

// 更新一条记录
func UpdateItem(c *gin.Context, model interface{}) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id <= 0 {
		response.Err(c, response.ErrMsg{
			Code: 707,
		})
		return
	}

	if ok := requests.Params(c, &model); !ok {
		return
	}

	if ok := requests.Check(c, model); !ok {
		return
	}

	if ok, errMsg := requests.ValidateRequest(model); !ok {
		response.Err(c, response.ErrMsg{
			Code:   414,
			Error:  errors.New("表单验证失败"),
			Result: errMsg,
		})
		return
	}

	if err := models.Update(id, model); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  701,
			Error: err,
		})
		return
	}

	response.Ok(c, nil)
}

// 删除一条记录
func DeleteItem(c *gin.Context, model interface{}) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	if id <= 0 {
		response.Err(c, response.ErrMsg{Code: 707})
	}

	if err := models.Find(id, model); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  700,
			Error: errors.New("删除的数据不存在"),
		})
		return
	}

	if err := models.Delete(id, model); err != nil {
		response.Err(c, response.ErrMsg{
			Code:  600,
			Error: err,
		})
		return
	} else {
		response.Ok(c, nil)
	}
}
