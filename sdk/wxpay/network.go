package wxpay

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var client *http.Client

func init() {
	client = &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout:     3 * time.Minute,
			TLSHandshakeTimeout: 10 * time.Second,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 10 * time.Minute,
			}).DialContext,
		},
	}
}

// 发送Get请求
func httpGet(url string) (body []byte, err error) {
	resp, err := client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// 发送Post请求
func httpPost(url string, body interface{}) (data []byte, err error) {
	bodyStr, err := json.Marshal(body)
	if err != nil {
		return
	}
	resp, err := client.Post(url, "application/json", bytes.NewReader(bodyStr))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	return
}

// 发送Post请求，参数是XML格式的字符串
func httpPostXml(url string, xmlBody string) (body []byte, err error) {
	resp, err := client.Post(url, "application/xml", strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}

// 发送带证书的Post请求，参数是XML格式的字符串
func httpPostXmlWithCert(url string, xmlBody string, client *http.Client) (body []byte, err error) {
	resp, err := client.Post(url, "application/xml", strings.NewReader(xmlBody))
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	return
}
