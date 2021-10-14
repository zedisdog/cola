package wechat

import (
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

const sendTemplate = "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token=%s"

func (s *Sdk) SendTemplate(data []byte, token string) (err error) {
	url := fmt.Sprintf(
		sendTemplate,
		token,
	)

	response, err := http.PostJSON(url, data)
	if err != nil {
		return err
	}
	err = hasErr(response)
	return
}
