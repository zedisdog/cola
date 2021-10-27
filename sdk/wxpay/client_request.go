package wxpay

import (
	"encoding/json"
)

func (c *Client) buildBody(bodyObj interface{}) (body map[string]interface{}, err error) {
	// 将bodyObj转换为map[string]interface{}类型
	bodyJson, _ := json.Marshal(bodyObj)
	body = make(map[string]interface{})
	_ = json.Unmarshal(bodyJson, &body)
	// 添加固定参数
	if c.isMch {
		body["mch_appid"] = c.config.AppId
		body["mchid"] = c.config.MchId
	} else {
		body["appid"] = c.config.AppId
		body["mch_id"] = c.config.MchId
	}
	if c.isFacilitator() {
		body["sub_appid"] = c.config.SubAppId
		body["sub_mch_id"] = c.config.SubMchId
	}
	nonceStr := GetRandomString(32)
	body["nonce_str"] = nonceStr
	// 生成签名
	signType, _ := body["sign_type"].(string)
	var sign string
	if c.isProd {
		sign = c.localSign(body, signType, c.apiKey)
	} else {
		body["sign_type"] = SignTypeMD5
		key, iErr := c.sandboxSign(nonceStr, SignTypeMD5)
		if err = iErr; iErr != nil {
			return
		}
		sign = c.localSign(body, SignTypeMD5, key)
	}
	body["sign"] = sign
	return
}

// 向微信发送请求
func (c *Client) doWeChat(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 转换参数
	body, err := c.buildBody(bodyObj)
	//fmt.Println(GenerateXml(body))
	if err != nil {
		return
	}
	// 发起请求
	bytes, err = httpPostXml(c.url(relativeUrl), GenerateXml(body))
	return
}

// 向微信发送带证书请求
func (c *Client) doWeChatWithCert(relativeUrl string, bodyObj interface{}) (bytes []byte, err error) {
	// 转换参数
	body, err := c.buildBody(bodyObj)
	if err != nil {
		return
	}
	// 设置证书和连接池
	if err = c.setCertData(c.certFilepath); err != nil {
		return
	}
	// 发起请求
	bytes, err = httpPostXmlWithCert(c.url(relativeUrl), GenerateXml(body), c.certClient)
	return
}
