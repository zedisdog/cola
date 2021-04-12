package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
	"github.com/zedisdog/cola/transport/http/response"
	"io"
	"net/http"
)

func ValidateJSON(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindJSON(request); err != nil {
		if e, ok := err.(validator.ValidationErrors); ok {
			c.AbortWithStatusJSON(422, response.NewValidateResponse(ParseValidateErrors(e)))
		} else if errors.Is(err, io.EOF) {
			c.AbortWithStatusJSON(400, response.Response{
				Msg: "body is empty",
			})
		} else {
			panic(err)
		}

		return err
	}

	return nil
}

func ValidateQuery(c *gin.Context, request interface{}) error {
	if err := c.ShouldBindQuery(request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response.NewValidateResponse(ParseValidateErrors(err.(validator.ValidationErrors))))
		return err
	}

	return nil
}

func ParseValidateErrors(errors validator.ValidationErrors) (message string, es map[string]string) {
	message = "the request is validate failed"
	es = make(map[string]string)
	for _, e := range errors {
		es[strcase.ToSnake(e.Field())] = e.(error).Error()
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
	Page    int `form:"page,default=1"`
	PerPage int `form:"per_page,default=20"`
}

func (p *PageQuery) Offset() int {
	return (p.Page - 1) * p.PerPage
}
