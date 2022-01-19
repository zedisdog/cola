package officialaccount

import (
	"github.com/zedisdog/cola/sdk/wechat/officialaccount/auth"
	"github.com/zedisdog/cola/sdk/wechat/officialaccount/sns"
)

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

func (c *Client) SNS() *sns.SNS {
	return sns.NewSNS()
}
