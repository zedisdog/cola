package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/cola/tools"
	"net/http"
)

func GenRoleMiddleware(roles ...string) func(*gin.Context) {
	return func(c *gin.Context) {
		if len(roles) < 1 {
			return
		}

		role, ok := c.Get("role")
		if !ok || !tools.InSlice(role, roles) {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"message": "没有权限",
			})
		}
	}
}
