package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/zedisdog/cola/auth"
	"net/http/httptest"
	"testing"
)

type myClaimsHavingRole struct {
	jwt.StandardClaims
	Role string
}

var key = "test"

func gen(claims jwt.Claims) *gin.Context {
	token, err := auth.GenerateToken([]byte(key), claims)
	if err != nil {
		panic(err)
	}

	request := httptest.NewRequest("get", "/test", nil)
	request.Header.Set("Authorization", fmt.Sprintf("Bearar %s", token))

	return &gin.Context{
		Request: request,
	}
}

func TestAuthMiddleware(t *testing.T) {
	var claims jwt.Claims = &myClaimsHavingRole{
		StandardClaims: jwt.StandardClaims{
			Id: "32185902184901",
		},
		Role: "admin",
	}
	cxt := gen(claims)
	f := GenAuthMiddleware(claims, "test", func(id interface{}) bool {
		return true
	})
	f(cxt)

	id, ok := cxt.Get("id")
	if !ok || id != "32185902184901" {
		t.Fatal("error")
	}

	role, ok := cxt.Get("role")
	if !ok || role != "admin" {
		t.Fatal("error2")
	}

	claims = &jwt.StandardClaims{
		Id: "32185902184901",
	}
	cxt = gen(claims)
	f = GenAuthMiddleware(claims, "test", func(id interface{}) bool {
		return true
	})
	f(cxt)
	_, ok = cxt.Get("role")
	if ok {
		t.Fatal("error4")
	}
}
