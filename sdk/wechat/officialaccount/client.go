package officialaccount

import "github.com/zedisdog/cola/sdk/wechat/officialaccount/auth"

func NewClient(AppID string, AppSecret string) *Client {
	return &Client{
		AppID:     AppID,
		AppSecret: AppSecret,
	}
}

type Client struct {
	AppID     string
	AppSecret string
}

func (c *Client) GetAuth(setters ...func(*auth.Auth)) *auth.Auth {
	return auth.NewAuth(c.AppID, c.AppSecret, setters...)
}
