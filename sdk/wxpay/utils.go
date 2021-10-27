package wxpay

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"time"
)

// 生成请求XML的Body体
func GenerateXml(data map[string]interface{}) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")
	for k, v := range data {
		buffer.WriteString(fmt.Sprintf("<%s><![CDATA[%v]]></%s>", k, v, k))
	}
	buffer.WriteString("</xml>")
	return buffer.String()
}

// 生成模型字符串
func JsonString(m interface{}) string {
	bs, _ := json.Marshal(m)
	return string(bs)
}

// 格式化时间，按照yyyyMMddHHmmss格式
func FormatDateTime(t time.Time) string {
	return t.Format("20060102150405")
}

// 对URL进行Encode编码
func EscapedPath(u string) (path string, err error) {
	uriObj, err := url.Parse(u)
	if err != nil {
		return
	}
	path = uriObj.EscapedPath()
	return
}

// 获取随机字符串
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	b := []byte(str)
	var result []byte
	loc, _ := time.LoadLocation("Asia/Shanghai")
	r := rand.New(rand.NewSource(time.Now().In(loc).UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

// 解密填充模式（去除补全码） PKCS7UnPadding
// 解密时，需要在最后面去掉加密时添加的填充byte
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unpadding := int(plainText[length-1])   // 找到Byte数组最后的填充byte
	return plainText[:(length - unpadding)] // 只截取返回有效数字内的byte数组
}

// 18位纯数字，以10、11、12、13、14、15开头
func IsValidAuthCode(authcode string) (ok bool) {
	pattern := "^1[0-5][0-9]{16}$"
	ok, _ = regexp.MatchString(pattern, authcode)
	return
}
