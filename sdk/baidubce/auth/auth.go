package auth

import (
	"encoding/json"
	"fmt"
	"github.com/uniplaces/carbon"
	"github.com/zedisdog/cola/sdk/baidubce/response"
	"io"
	"net/http"
	"net/url"
	"time"
)

func NewAuth(clientId string, clientSecret string, host string) *Auth {
	return &Auth{
		clientId:     clientId,
		clientSecret: clientSecret,
		host:         host,
	}
}

type Auth struct {
	host         string
	clientId     string
	clientSecret string
	expires      *time.Time
	accessToken  string
}

func (a Auth) GetAccessToken() (token string, err error) {
	refresh := false
	if a.expires == nil {
		refresh = true
	}
	if !refresh {
		var now *carbon.Carbon
		now, err = carbon.NowInLocation("Asia/Shanghai")
		if err != nil {
			return
		}
		isExpires := carbon.NewCarbon(*a.expires).Lt(now.SubDay())
		if !isExpires {
			refresh = true
		}
	}

	if refresh {
		err = a.doGetAccessToken()
		if err != nil {
			return
		}
	}

	return a.accessToken, nil
}

func (a *Auth) doGetAccessToken() (err error) {
	u := url.URL{
		Scheme: "https",
		Host:   a.host,
		Path:   "oauth/2.0/token",
		RawQuery: fmt.Sprintf(
			"grant_type=%s&client_id=%s&client_secret=%s",
			"client_credentials",
			a.clientId,
			a.clientSecret,
		),
	}
	r, err := http.Get(u.String())
	if err != nil {
		return
	}
	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}
	if r.StatusCode >= 400 {
		return response.ParseError(r)
	}
	var res authSuccess
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	a.accessToken = res.AccessToken

	now, err := carbon.NowInLocation("Asia/Shanghai")
	if err != nil {
		return
	}
	a.expires = &now.AddSeconds(res.ExpiresIn).Time
	return
}

type authSuccess struct {
	RefreshToken  string `json:"refresh_token"`
	ExpiresIn     int    `json:"expires_in"`
	Scope         string `json:"scope"`
	SessionKey    string `json:"session_key"`
	AccessToken   string `json:"access_token"`
	SessionSecret string `json:"session_secret"`
}
