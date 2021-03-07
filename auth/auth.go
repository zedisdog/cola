package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

var (
	TokenIsInvalid = errors.New("token is invalid")
)

func GenerateToken(key []byte, claims jwt.Claims) (string, error) {
	if claims == nil {
		panic("clamis is required")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return ss, nil
}

func ParseToken(token string, key []byte, claims jwt.Claims) (jwt.Claims, error) {
	t, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		return nil, err
	}
	if t.Valid {
		return t.Claims, nil
	} else {
		return nil, TokenIsInvalid
	}
}
