package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/uniplaces/carbon"
	"testing"
)

func TestExpire(t *testing.T) {
	key := []byte("123")
	claims := jwt.StandardClaims{
		ExpiresAt: carbon.Now().AddMinute().Unix(),
	}
	token, err := GenerateToken(key, claims)
	if err != nil {
		t.Fatal(err)
	}
	err = Parse(token, key, &claims)
	if err != nil {
		t.Fatal(err)
	}

	claims = jwt.StandardClaims{
		ExpiresAt: carbon.Now().SubSecond().Unix(),
	}
	token, err = GenerateToken(key, claims)
	if err != nil {
		t.Fatal(err)
	}
	err = Parse(token, key, &claims)
	if err == nil {
		t.Fatal("token is still active")
	}
}
