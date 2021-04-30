package response

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/cola/e"
	"github.com/zedisdog/cola/errx"
	"gorm.io/gorm"
	"math"
	"net/http"
)

type Meta struct {
	CurrentPage uint `json:"current_page"`
	Total       uint `json:"total"`
	LastPage    uint `json:"last_page"`
	PerPage     uint `json:"per_page"`
}

type Response struct {
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
	Meta *Meta       `json:"meta"`
}

func NewValidateResponse(message string, data interface{}) *Response {
	return &Response{
		Data: data,
		Msg:  message,
	}
}

// Deprecated: func is too many
func NewErrorResponse(err error) *Response {
	return &Response{
		Msg: err.Error(),
	}
}

func NewResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

func NewPageResponse(data interface{}, total int, page int, perPage int) *Response {
	return &Response{
		Meta: &Meta{
			CurrentPage: uint(page),
			Total:       uint(total),
			LastPage:    uint(math.Ceil(float64(total) / float64(perPage))),
			PerPage:     uint(perPage),
		},
		Data: data,
	}
}

func NewPageResponseWithMeta(data interface{}, meta *Meta) *Response {
	return &Response{
		Meta: meta,
		Data: data,
	}
}

// Error 返回错误响应 p1 错误 p2 status code
func Error(c *gin.Context, params ...interface{}) {
	if len(params) == 0 {
		panic("need at least err")
	}
	err := params[0].(error)
	var code int
	if len(params) == 2 {
		code = params[1].(int)
	} else {
		if errors.Is(err, e.NotFoundError) || errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		} else if errors.Is(err, e.ConflictError) {
			code = http.StatusConflict
		} else if er, ok := err.(*HttpError); ok {
			code = er.StatusCode
		} else if er, ok := err.(*errx.HttpError); ok {
			code = er.StatusCode
		} else {
			code = http.StatusInternalServerError
		}
	}

	res := &Response{Msg: err.Error()}
	if er, ok := err.(*errx.HttpError); ok && code == http.StatusTeapot {
		res.Data = er.Data
	}

	c.AbortWithStatusJSON(code, res)
}

// Success params[0]: data
//         params[1]: status code
func Success(c *gin.Context, params ...interface{}) {
	if len(params) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	var code int
	if len(params) == 2 {
		code = params[1].(int)
	} else {
		code = http.StatusOK
	}
	var response *Response
	if list, ok := params[0].(*ListByPage); ok {
		response = NewPageResponse(list.Data, list.Total, list.Page, list.PerPage)
	} else {
		response = NewResponse(params[0])
	}
	c.JSON(code, response)
}

func Pagination(c *gin.Context, data interface{}, total int, page int, perPage int) {
	response := NewPageResponse(data, total, page, perPage)
	c.JSON(http.StatusOK, response)
}

// Deprecated: use Pagination instead
func ReturnPagination(c *gin.Context, data interface{}, total int, page int, perPage int) {
	response := NewPageResponse(data, total, page, perPage)
	c.JSON(http.StatusOK, response)
}

type ListByPage struct {
	Data     interface{}
	Total    int
	Page     int
	PerPage  int
	LastPage int
}

func NewListByPage(data interface{}, total int, page int, perPage int) *ListByPage {
	return &ListByPage{
		Data:     data,
		Total:    total,
		Page:     page,
		PerPage:  perPage,
		LastPage: int(math.Ceil(float64(total) / float64(perPage))),
	}
}

func Return(cxt *gin.Context, data interface{}, err error) {
	if err != nil {
		Error(cxt, err)
		return
	}
	Success(cxt, data)
}
