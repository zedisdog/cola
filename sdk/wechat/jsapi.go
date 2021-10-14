package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/zedisdog/cola/transport/http"
)

const getTicketUrl = "https://api.weixin.qq.com/cgi-bin/ticket/getticket?access_token=%s&type=jsapi"

//GetTicket get js ticket for wechat
func (s *Sdk) GetTicket(accountToken string) (ticket Ticket, err error) {
	url := fmt.Sprintf(getTicketUrl, accountToken)
	response, err := http.GetJSON(url)
	if err != nil {
		return
	}
	if err = hasErr(response); err != nil {
		return
	}

	err = json.Unmarshal(response, &ticket)
	return
}

type Ticket struct {
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}
