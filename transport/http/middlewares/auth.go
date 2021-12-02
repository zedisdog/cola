package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zedisdog/cola/auth"
	"net/http"
	"reflect"
	"strings"
)

//GenAuthMiddleware generate an auth middleware
//	claims: the customer Claims object
//	key: the key for validate sign
//	isUserExists: a function to determine if account is exists
func GenAuthMiddleware(claims jwt.Claims, key string, isUserExists func(id interface{}) bool) func(*gin.Context) {
	return func(c *gin.Context) {
		var token string
		if c.Request.Header.Get("Authorization") != "" {
			arr := strings.Split(c.Request.Header.Get("Authorization"), " ")
			if len(arr) < 2 {
				c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问1"})
				return
			}
			token = arr[1]
		} else if c.Query("token") != "" {
			token = c.Query("token")
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问2"})
			return
		}

		err := auth.Parse(token, []byte(key), claims)
		if err != nil {
			fmt.Printf("%+v\n", err)
			fmt.Printf("%+v\n", token)
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问3"})
			return
		}

		valueOfClaims := reflect.ValueOf(claims)

		// if claims has property 'Role', set it into context
		roleField := valueOfClaims.Elem().FieldByName("Role")
		if roleField.IsValid() && roleField.String() != "" {
			c.Set("role", roleField.String())
		}

		// use jti to keep the identification of account, jti in jwt.StandardClaims is Id
		id := valueOfClaims.Elem().FieldByName("Id").String()
		if !isUserExists(id) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]string{"message": "未授权的访问4"})
			return
		}
		c.Set("id", id)

		c.Set("claims", claims)

		c.Next()
	}
}
