package auth

import (
	"fmt"
	"github.com/zedisdog/cola/sdk/wechat/util"
	"github.com/zedisdog/cola/transport/http"
	"net/url"
	"strings"
)

const (
	redirectUrl      = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect"
	exchangeUrl      = "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code"
	refreshUrl       = "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s"
	userInfoUrl      = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
	validateTokenUrl = "https://api.weixin.qq.com/sns/auth?access_token=%s&openid=%s"
	SNSAPI_BASE      = "snsapi_base"
	SNSAPI_USERINFO  = "snsapi_userinfo"
)

func WithLang(lang string) func(*Auth) {
	return func(auth *Auth) {
		auth.Lang = lang
	}
}

func NewAuth(AppID string, AppSecret string, setters ...func(*Auth)) *Auth {
	auth := &Auth{
		AppID:     AppID,
		AppSecret: AppSecret,
		Lang:      "zh_CN",
	}
	for _, set := range setters {
		set(auth)
	}
	return auth
}

type Auth struct {
	AppID     string
	AppSecret string
	Lang      string
}

func (c *Auth) GenRedirectUrl(redirectURI string, scope string, states ...string) string {
	return fmt.Sprintf(
		redirectUrl,
		c.AppID,
		redirectURI,
		scope,
		strings.Join(states, ","),
	)
}

//AuthResponse 微信网页授权回调的code和state
type AuthResponse struct {
	Code  string
	State []string
}

func NewAuthResponse(query string) (response *AuthResponse, err error) {
	values, err := url.ParseQuery(query)
	if err != nil {
		return
	}

	response = new(AuthResponse)
	response.State = strings.Split(values.Get("state"), ",")
	response.Code = values.Get("code")

	return
}

//AccessToken struct of response
//  {
//    "access_token":"ACCESS_TOKEN",
//    "expires_in":7200,
//    "refresh_token":"REFRESH_TOKEN",
//    "openid":"OPENID",
//    "scope":"SCOPE"
//  }
type AccessToken struct {
	util.Error
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

//ExToken exchange access token of user by code.
func (a *Auth) ExToken(code string) (token AccessToken, err error) {
	u := fmt.Sprintf(
		exchangeUrl,
		a.AppID,
		a.AppSecret,
		code,
	)
	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &token)
	return
}

func (a *Auth) RefreshToken(refreshToken string) (token AccessToken, err error) {
	u := fmt.Sprintf(
		refreshUrl,
		a.AppID,
		refreshToken,
	)

	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &token)
	return
}

//UserInfo is struct of user info
//  {
//    "openid": "OPENID",
//    "nickname": NICKNAME,
//    "sex": 1,
//    "province":"PROVINCE",
//    "city":"CITY",
//    "country":"COUNTRY",
//    "headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
//    "privilege":["PRIVILEGE1", "PRIVILEGE2"],
//    "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
//  }
type UserInfo struct {
	util.Error
	Openid     string   `json:"openid"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	UnionID    string   `json:"union_id"`
}

//UserInfo get user info.
func (a *Auth) UserInfo(accessToken string, openID string) (info UserInfo, err error) {
	u := fmt.Sprintf(
		userInfoUrl,
		accessToken,
		openID,
		a.Lang,
	)

	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &info)
	return
}

//ValidateAccessToken validates if is valid which given accessToken of openID.
func (a *Auth) ValidateAccessToken(accessToken string, openID string) bool {
	var errMsg util.Error
	u := fmt.Sprintf(
		validateTokenUrl,
		accessToken,
		openID,
	)

	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &errMsg)
	if err != nil {
		println(err.Error())
		return false
	}

	return errMsg.Errcode == 0
}
