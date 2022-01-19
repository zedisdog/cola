package sns

import (
	"fmt"
	"github.com/zedisdog/cola/sdk/wechat/util"
	"github.com/zedisdog/cola/transport/http"
)

const (
	userInfoUrl = "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=%s"
)

type UserInfoRequest struct {
	AccessToken string
	OpenID      string
	Lang        string
}

func NewUserInfoRequest(accessToken, OpenID string, lang ...string) *UserInfoRequest {
	r := &UserInfoRequest{
		AccessToken: accessToken,
		OpenID:      OpenID,
		Lang:        "zh_CN",
	}
	if len(lang) > 0 {
		r.Lang = lang[0]
	}
	return r
}

func NewSNS() *SNS {
	return &SNS{}
}

type SNS struct{}

//UserInfo {
//  "openid": "OPENID",
//  "nickname": NICKNAME,
//  "sex": 1,
//  "province":"PROVINCE",
//  "city":"CITY",
//  "country":"COUNTRY",
//  "headimgurl":"https://thirdwx.qlogo.cn/mmopen/g3MonUZtNHkdmzicIlibx6iaFqAc56vxLSUfpb6n5WKSYVY0ChQKkiaJSgQ1dZuTOgvLLrhJbERQQ4eMsv84eavHiaiceqxibJxCfHe/46",
//  "privilege":[ "PRIVILEGE1" "PRIVILEGE2"     ],
//  "unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
//}
type UserInfo struct {
	util.Error
	OpenID     string   `json:"open_id"`
	Nickname   string   `json:"nickname"`
	Sex        int      `json:"sex"`
	Province   string   `json:"province"`
	City       string   `json:"city"`
	Country    string   `json:"country"`
	Headimgurl string   `json:"headimgurl"`
	Privilege  []string `json:"privilege"`
	Unionid    string   `json:"unionid"`
}

//GetUserInfo get user info by access_token.
func (s *SNS) GetUserInfo(request *UserInfoRequest) (userInfo UserInfo, err error) {
	u := fmt.Sprintf(
		userInfoUrl,
		request.AccessToken,
		request.OpenID,
		request.Lang,
	)
	response, err := http.GetJSON(u)
	if err != nil {
		return
	}
	err = util.ParseResponse(response, err, &userInfo)
	return
}
