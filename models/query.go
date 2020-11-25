package models

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	PAGESIZE = 20 // 每页默认数量
	PAGE     = 1  // 当前页码
)

// 查询实例
type Query struct {
	Ctx         *gin.Context // gin context
	DB          *gorm.DB     // DB 实例
	SelectField []string     // 指定查询字段
	FilterField interface{}  // 可搜索字段
	Pagination  Pagination   // 分页
	count       int64        // 结果统计
}

// 分页信息
type Pagination struct {
	TotalRecord int64 `json:"total_record"`
	CurrentPage int   `json:"current_page"`
	PageSize    int   `json:"page_size"`
	TotalPages  int   `json:"total_pages"`
	HasNext     bool  `json:"has_next"`
}

// 查询条件
type condition struct {
	Query string
	Args  interface{}
	IsOr  bool
}

func (q *Query) Fetch(model interface{}) (err error) {
	q.preFetch()
	var res *gorm.DB

	offset := q.Pagination.PageSize * (q.Pagination.CurrentPage - 1)
	limit := q.Pagination.PageSize
	fmt.Println(q.SelectField)
	q.DB.Model(model).Select(q.SelectField).Count(&q.count)
	res = q.DB.Model(model).Select(q.SelectField).Offset(offset).Limit(limit).Find(model)

	q.Pagination.TotalRecord = q.count
	q.Pagination.TotalPages = int(math.Ceil(float64(q.count) / float64(q.Pagination.PageSize)))
	q.Pagination.HasNext = q.Pagination.TotalPages > q.Pagination.CurrentPage

	return res.Error
}

// 查询条件
func (q *Query) preFetch() {
	q.applyWhere()
	q.applyPage()
	q.applyOrder()
}

// 解析查询参数
// ?order=id:desc,age:asc
func (q *Query) applyOrder() {
	sorter, has := q.Ctx.GetQuery("order")
	if has && sorter != "" {
		orders := q.parseOrder(sorter)

		for _, item := range orders {
			q.DB = q.DB.Order(strings.Join(item, " "))
		}
	} else { //添加一个默认的排序，防止分页时记录可能会重复出现的问题
		q.DB = q.DB.Order("id desc")
	}
}

// format order query
func (q Query) parseOrder(sorter string) (sorterSlice [][]string) {
	sorterMap := strings.Split(sorter, ",")
	if len(sorterMap) > 0 {
		for _, item := range sorterMap {
			itemMap := strings.SplitN(item, ":", 2)
			itemLen := len(itemMap)
			if itemLen > 0 {
				s := []string{
					itemMap[0],
					"asc",
				}

				if itemLen > 1 {
					sortedBy := strings.ToLower(itemMap[1])
					if sortedBy == "desc" || sortedBy == "asc" {
						s[1] = sortedBy
					}
				}

				sorterSlice = append(sorterSlice, s)
			}
		}
	}
	return
}

// 解析分页参数
// ?page=2&page_size=10
func (q *Query) applyPage() {
	if q.Pagination.CurrentPage == 0 {
		page, hasPage := q.Ctx.GetQuery("page")

		if hasPage && page != "" {
			p, e := strconv.Atoi(page)
			if e == nil {
				q.Pagination.CurrentPage = p
			}
		}

		if q.Pagination.CurrentPage < 1 {
			q.Pagination.CurrentPage = PAGE
		}
	}

	if q.Pagination.PageSize == 0 {
		pageSize, hasSize := q.Ctx.GetQuery("page_size")
		if hasSize && pageSize != "" {
			s, e := strconv.Atoi(pageSize)
			if e == nil {
				q.Pagination.PageSize = s
			}
		}
		if q.Pagination.PageSize < 1 { //默认10条
			q.Pagination.PageSize = PAGESIZE
		}
	}
}

// 查询条件
func (q *Query) applyWhere() {
	whereSlice := q.parseWhere()
	if whereSlice == nil {
		return
	}
	for _, cond := range whereSlice {
		if cond.IsOr {
			q.DB = q.DB.Or(cond.Query, cond.Args)
		} else {
			q.DB = q.DB.Where(cond.Query, cond.Args)
		}
	}
}

// 格式他查询条件
func (q Query) parseWhere() (cond []condition) {
	if q.FilterField == nil {
		return
	}
	s := reflect.TypeOf(q.FilterField)
	for i := 0; i < s.NumField(); i++ {
		key := s.Field(i).Tag.Get("form")
		args, ok := q.Ctx.GetQuery(key)
		if ok {
			field := s.Field(i).Tag.Get("field")
			operator := s.Field(i).Tag.Get("operator")
			isOr, _ := strconv.ParseBool(s.Field(i).Tag.Get("is_or"))
			var where string
			switch operator {
			case "in":
				where = field + " IN (?)"
			case "like":
				where = field + " LIKE ?"
			default:
				where = field + " " + operator + " ?"
			}
			cond = append(cond, condition{
				Query: where,
				Args:  args,
				IsOr:  isOr,
			})
		}
	}
	return
}
