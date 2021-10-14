package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

const (
	menuCreateURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s"
	//menuGetURL               = "https://api.weixin.qq.com/cgi-bin/menu/get"
	menuSelfMenuInfoURL = "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token=%s"
)

//Menu 菜单结构体
type Menu struct {
	IsMenuOpen   int64 `json:"is_menu_open"`
	SelfMenuInfo struct {
		Button []interface{} `json:"button"`
	} `json:"selfmenu_info"`
}

//GetMenu 获取菜单内容
func (s *Sdk) GetMenu(token string) (menu Menu, err error) {
	url := fmt.Sprintf(
		menuSelfMenuInfoURL,
		token,
	)

	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &menu)
	return
}

//CreateMenu 创建菜单
func (s *Sdk) CreateMenu(menu []byte, token string) (err error) {
	url := fmt.Sprintf(
		menuCreateURL,
		token,
	)

	response, err := http.PostJSON(url, menu)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}
	return
}
