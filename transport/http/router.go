package http

import (
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/cola/transport/http/middlewares"
)

type RouteSetter func(*gin.RouterGroup)

func NewRouter(setters ...RouteSetter) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.Cros)
	r.Static("public", "storage/public")
	api := r.Group("api")
	for _, setter := range setters {
		setter(api)
	}
	return r
}
