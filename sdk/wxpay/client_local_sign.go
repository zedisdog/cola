package wxpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

// 本地通过支付参数计算签名值
// 生成算法：https://pay.weixin.qq.com/wiki/doc/api/micropay.php?chapter=4_3
func (c *Client) localSign(body map[string]interface{}, signType string, apiKey string) string {
	signStr := c.sortSignParams(body, apiKey)
	//fmt.Println(signStr)
	var hashSign []byte
	if signType == SignTypeHmacSHA256 {
		hash := hmac.New(sha256.New, []byte(apiKey))
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	} else {
		hash := md5.New()
		hash.Write([]byte(signStr))
		hashSign = hash.Sum(nil)
	}
	return strings.ToUpper(hex.EncodeToString(hashSign))
}

// 获取根据Key排序后的请求参数字符串
func (c *Client) sortSignParams(body map[string]interface{}, apiKey string) string {
	keyList := make([]string, 0)
	for k := range body {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	buffer := new(bytes.Buffer)
	for _, k := range keyList {
		s := fmt.Sprintf("%s=%s&", k, fmt.Sprintf("%v", body[k]))
		buffer.WriteString(s)
	}
	buffer.WriteString(fmt.Sprintf("key=%s", apiKey))
	return buffer.String()
}
