package util

//Error is the wechat error
type Error struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}
