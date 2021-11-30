package util

type WechatError interface {
	ErrCode() int
	ErrMsg() string
}

//Error is the wechat error
type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

func (e Error) ErrCode() int {
	return e.Errcode
}

func (e Error) ErrMsg() string {
	return e.Errmsg
}
