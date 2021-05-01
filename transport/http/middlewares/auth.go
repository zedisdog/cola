package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/cola/auth"
	"net/http"
	"reflect"
	"strings"
)

func GenAuthMiddleware(claims jwt.Claims, key string, isUserExists func(id interface{}) bool) func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Header.Get("Authorization") != "" {
			arr := strings.Split(c.Request.Header.Get("Authorization"), " ")
			if len(arr) < 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问1"})
				return
			}
			token := arr[1]
			err := auth.Parse(token, []byte(key), claims)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问2"})
				return
			}

			cv := reflect.ValueOf(claims)
			id := cv.Elem().FieldByName("Id").String()
			if !isUserExists(id) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问3"})
				return
			}
			c.Set("id", id)

			roleField := cv.Elem().FieldByName("Role")
			if roleField.IsValid() && roleField.String() != "" {
				c.Set("role", roleField.String())
			}
		} else if c.Query("token") != "" {
			err := auth.Parse(c.Query("token"), []byte(key), claims)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问2"})
				return
			}

			cv := reflect.ValueOf(claims)
			id := cv.Elem().FieldByName("Id").String()
			if !isUserExists(id) {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问3"})
				return
			}
			c.Set("id", id)

			roleField := cv.Elem().FieldByName("Role")
			if roleField.IsValid() && roleField.String() != "" {
				c.Set("role", roleField.String())
			}
		}

		c.Next()
	}
}
