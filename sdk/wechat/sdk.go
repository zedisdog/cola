package wechat

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Sdk struct {
	AppID         string
	AppSecret     string
	Token         string
	AesKey        string
	MiniAppID     string
	MiniAppSecret string
	log           *logrus.Logger
}

type SdkOptions struct {
	AppID     string
	AppSecret string
	Token     string
	AesKey    string
}

//ErrorResponse {"errcode":41001,"errmsg":"access_token missing hint: [aMOoRa04274693!]"}
type ErrorResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func NewOptions(v *viper.Viper) *SdkOptions {
	return &SdkOptions{
		AppID:     v.GetString("wx.wechat.appId"),
		AppSecret: v.GetString("wx.wechat.secret"),
		Token:     v.GetString("wx.wechat.token"),
		AesKey:    v.GetString("wx.wechat.aesKey"),
	}
}

func New(options *SdkOptions, log *logrus.Logger) *Sdk {
	return &Sdk{
		AppID:     options.AppID,
		AppSecret: options.AppSecret,
		Token:     options.Token,
		AesKey:    options.AesKey,
		log:       log,
	}
}

func hasErr(response []byte) error {
	//log.Log.Infoln(string(response))
	errResponse := &ErrorResponse{}
	_ = json.Unmarshal(response, errResponse)
	//log.Log.Infoln(errResponse.Errcode)
	if errResponse.Errcode != 0 {
		return errors.New(string(response))
	}
	return nil
}
