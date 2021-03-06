package middlewares

import (
	"github.com/gin-gonic/gin"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache")
	c.Header("Pragma", "no-cache")
	c.Next()
}
