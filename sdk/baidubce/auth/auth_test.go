package auth

import (
	"fmt"
	"net/url"
	"testing"
)

func TestNormal(t *testing.T) {
	u := url.URL{
		Scheme: "https",
		Host:   "localhost", //todo: 出错先看host
		Path:   "oauth/2.0/token",
		RawQuery: fmt.Sprintf(
			"grant_type=%s&client_id=%s&client_secret=%s",
			"client_credentials",
			"123",
			"456",
		),
	}
	println(u.String())
}
