package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

const (
	//UserInfoURL API接口地址
	UserInfoURL = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN" // ?access_token=ACCESS_TOKEN&openid=OPENID&lang=zh_CN
)

/*
{
    "subscribe": 1,
    "openid": "o6_bmjrPTlm6_2sgVt7hMZOPfL2M",
    "nickname": "Band",
    "sex": 1,
    "language": "zh_CN",
    "city": "广州",
    "province": "广东",
    "country": "中国",
    "headimgurl":"http://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/0",
    "subscribe_time": 1382694957,
    "unionid": " o6_bmasdasdsad6_2sgVt7hMZOPfL"
    "remark": "",
    "groupid": 0,
    "tagid_list":[128,2],
    "subscribe_scene": "ADD_SCENE_QR_CODE",
    "qr_scene": 98765,
    "qr_scene_str": ""
}
*/

//UserInfo 微信用户信息
type UserInfo struct {
	Subscribe      int    `json:"subscribe"`
	OpenID         string `json:"-"`
	Nickname       string `json:"nickname"`
	Sex            int    `json:"sex"`
	Language       string `json:"language"`
	City           string `json:"city"`
	Province       string `json:"province"`
	Country        string `json:"country"`
	Headimgurl     string `json:"headimgurl"`
	SubscribeTime  int    `json:"subscribe_time"`
	Unionid        string `json:"unionid"`
	Remark         string `json:"remark"`
	GroupID        int    `json:"groupid"`
	SubscribeScene string `json:"subscribe_scene"`
	QrScene        int    `json:"qr_scene"`
	QrSceneStr     string `json:"qr_scene_str"`
}

//GetUserInfo 获取微信用户信息
func (s *Sdk) GetUserInfo(openid string, token string) (userInfo UserInfo, err error) {
	url := fmt.Sprintf(
		UserInfoURL,
		token,
		openid,
	)

	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &userInfo)
	return
}
