package util

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/errx"
)

func ParseResponse(response []byte, err error, data interface{}) error {
	if err != nil {
		return errx.Wrap(err, "get user info error")
	}

	err = json.Unmarshal(response, data)
	if err != nil {
		return errx.Wrap(err, "parse response error")
	}

	if data.(Error).Errcode != 0 {
		err = errx.New(fmt.Sprintf("refresh token error: errcode=%d errmsg=%s", data.(Error).Errcode, data.(Error).Errmsg))
	}
	return err
}
