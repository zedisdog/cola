package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

//GET https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
const authCode2SessionURL = "https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code"

//AuthInfo 结果
type AuthInfo struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

// Auth 小程序登录
func (s *Sdk) Auth(jsCode string) (info AuthInfo, err error) {
	url := fmt.Sprintf(
		authCode2SessionURL,
		s.MiniAppID,
		s.MiniAppSecret,
		jsCode,
	)

	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &info)
	return
}
