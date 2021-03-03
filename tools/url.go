package tools

import (
	"fmt"
	"net/url"
)

func EncodeQuery(u string) string {
	tmp, err := url.Parse(u)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s://%s@%s%s?%s", tmp.Scheme, tmp.User, tmp.Host, tmp.Path, tmp.Query().Encode())
}
