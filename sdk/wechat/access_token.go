package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

const (
	accessTokenURL      = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	code2AccessTokenURL = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code" //?appid=APPID&secret=SECRET&code=CODE&grant_type=authorization_code
)

// AccessToken get the access token of wechat app
func (s *Sdk) AccessToken() (token AppToken, err error) {
	url := fmt.Sprintf(
		accessTokenURL,
		s.AppID,
		s.AppSecret,
	)

	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &token)
	return
}

//Code2AccessToken exchange for access token of wechat user using code
func (s *Sdk) Code2AccessToken(code string) (token UserToken, err error) {
	url := fmt.Sprintf(
		code2AccessTokenURL,
		s.AppID,
		s.AppSecret,
		code,
	)

	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &token)
	return
}

//AppToken app token of wechat
type AppToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//UserToken user token of wechat
type UserToken struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int64  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	OpenID       string `json:"openid"`
	Scope        string `json:"scope"`
}
