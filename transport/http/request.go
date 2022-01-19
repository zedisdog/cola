package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/zedisdog/cola/errx"
	"github.com/zedisdog/cola/i18n"
	"io"
)

func ValidateJSON(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindJSON(request); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			return errx.NewHttpErrorUnprocessableEntityWithDetail(ParseValidateErrors(e, request))
		} else if errors.Is(err, io.EOF) {
			return errx.NewHttpErrorBadRequest(i18n.Trans(i18n.EMPTY_BODY))
		} else {
			panic(err)
		}

		return err
	}

	return nil
}

type CanGetError interface {
	GetError(structField string, tag string) string
}

func ValidateQuery(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindQuery(request); err != nil {
		return errx.NewHttpErrorUnprocessableEntityWithDetail(ParseValidateErrors(err.(validator.ValidationErrors), request))
	}

	return nil
}

func ParseValidateErrors(errors validator.ValidationErrors, request interface{}) (message string, es map[string]string) {
	es = make(map[string]string)
	for _, e := range errors {
		key := strcase.ToSnake(e.StructField())
		if _, ok := es[key]; !ok {
			var m string
			if r, ok := request.(CanGetError); ok {
				m = r.GetError(e.StructField(), e.Tag())
			}
			if m == "" {
				m = e.(error).Error()
			}
			es[key] = m
			if message == "" {
				message = m
			}
		}
	}

	return
}

type FetchAndPageQuery struct {
	FetchQuery
	PageQuery
}

type FetchQuery struct {
	Fetch bool `form:"fetch,default=false"`
}

type PageQuery struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=15"`
}

func (p *PageQuery) Offset() int {
	return (p.Page - 1) * p.Size
}
