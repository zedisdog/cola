package auth

import (
	"fmt"
	"github.com/zedisdog/cola/cache"
	"github.com/zedisdog/cola/errx"
	"github.com/zedisdog/cola/sdk/wechat/util"
	"github.com/zedisdog/cola/transport/http"
	"time"
)

const (
	tokenUrl    = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	jsTicketUrl = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"
)

type Auth struct {
	AppID     string
	AppSecret string
}

func NewAuth(AppID string, AppSecret string) *Auth {
	auth := &Auth{
		AppID:     AppID,
		AppSecret: AppSecret,
	}
	return auth
}

func (a *Auth) AccessToken() string {
	key := fmt.Sprintf("access_token|%s", a.AppID)
	token, ok := cache.Get(key)
	if !ok {
		var err error
		token, err = a.GetAccessToken()
		if err != nil {
			fmt.Printf("get access token error: %s", errx.Wrap(err, "error"))
			return ""
		}
		cache.PutWithExpire(key, token, time.Now().Add(time.Duration(token.(Token).ExpiresIn-60)*time.Second).Unix())
	}

	return token.(Token).AccessToken
}

//Token {"access_token":"ACCESS_TOKEN","expires_in":7200}
type Token struct {
	util.Error
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func (a *Auth) GetAccessToken() (token Token, err error) {
	u := fmt.Sprintf(
		tokenUrl,
		a.AppID,
		a.AppSecret,
	)
	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &token)
	return
}

func (a *Auth) JsTicket() string {
	key := fmt.Sprintf("js_ticket|%s", a.AppID)
	ticket, ok := cache.Get(key)
	if !ok {
		var err error
		ticket, err = a.GetJsTicket()
		if err != nil {
			fmt.Printf("get js ticket error: %s", errx.Wrap(err, "error"))
			return ""
		}
		cache.PutWithExpire(key, ticket, time.Now().Add(time.Duration(ticket.(JsTicket).ExpiresIn-60)*time.Second).Unix())
	}

	return ticket.(JsTicket).Ticket
}

// JsTicket
// {
//  "errcode":0,
//  "errmsg":"ok",
//  "ticket":"bxLdikRXVbTPdHSM05e5u5sUoXNKd8-41ZO3MhKoyN5OfkWITDGgnr2fwJ0m9E8NYzWKVZvdVtaUgWvsdshFKA",
//  "expires_in":7200
// }
type JsTicket struct {
	util.Error
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

func (a *Auth) GetJsTicket() (ticket JsTicket, err error) {
	u := fmt.Sprintf(
		jsTicketUrl,
		a.AccessToken(),
	)
	response, err := http.GetJSON(u)
	err = util.ParseResponse(response, err, &ticket)
	return
}
