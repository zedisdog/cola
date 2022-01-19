package response

import (
	"errors"
	"github.com/gin-gonic/gin"
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

// Error 返回错误响应 p1 错误 p2 status code
func Error(c *gin.Context, errAndStatus ...interface{}) {
	if len(errAndStatus) == 0 {
		panic("need at least err")
	}

	err := errAndStatus[0].(error)
	res := &Response{Msg: err.Error()}

	var code int
	if len(errAndStatus) == 2 {
		code = errAndStatus[1].(int)
	} else {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			code = http.StatusNotFound
		} else if er, ok := err.(*errx.HttpError); ok {
			code = er.StatusCode
			res.Data = er.Data
		} else {
			code = http.StatusInternalServerError
		}
	}

	c.JSON(code, res)
}

// Success params[0]: data
//         params[1]: status code
func Success(c *gin.Context, params ...interface{}) {
	var (
		code     int
		response *Response
	)

	switch len(params) {
	case 0:
		code = http.StatusNoContent
	case 1:
		response = &Response{
			Data: params[0],
		}
		code = http.StatusOK
	case 2:
		response = &Response{
			Data: params[0],
		}
		code = params[1].(int)
	}

	c.JSON(code, response)
}

func Pagination(c *gin.Context, data interface{}, total int, page int, perPage int) {
	resp := &Response{
		Meta: &Meta{
			CurrentPage: uint(page),
			Total:       uint(total),
			LastPage:    uint(math.Ceil(float64(total) / float64(perPage))),
			PerPage:     uint(perPage),
		},
		Data: data,
	}
	if resp.Meta.CurrentPage == 0 {
		resp.Meta.CurrentPage = 1
	}
	if resp.Meta.LastPage == 0 {
		resp.Meta.LastPage = 1
	}
	c.JSON(http.StatusOK, resp)
}

func Return(cxt *gin.Context, data interface{}, err error) {
	if err != nil {
		Error(cxt, err)
		return
	}
	Success(cxt, data)
}
