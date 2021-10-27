package wxpay

import (
	"encoding/xml"
	"errors"
)

// 获取沙盒签名Key的返回值
type getSignKeyResponse struct {
	ResponseModel
	Retcode        string `xml:"retcode"`
	MchId          string `xml:"mch_id"`
	SandboxSignkey string `xml:"sandbox_signkey"`
}

// 获取沙盒的签名
func (c *Client) sandboxSign(nonceStr string, signType string) (key string, err error) {
	body := make(BodyMap)
	body["mch_id"] = c.config.MchId
	body["nonce_str"] = nonceStr
	// 计算沙箱参数Sign
	sanboxSign := c.localSign(body, signType, c.apiKey)
	// 沙箱环境：获取key后，重新计算Sign
	key, err = c.getSandBoxSignKey(nonceStr, sanboxSign)
	return
}

// 调用微信提供的接口获取SandboxSignkey
func (c *Client) getSandBoxSignKey(nonceStr string, sign string) (key string, err error) {
	params := make(map[string]interface{})
	params["mch_id"] = c.config.MchId
	params["nonce_str"] = nonceStr
	params["sign"] = sign
	paramXml := GenerateXml(params)
	bytes, err := httpPostXml(baseUrlSandbox+"pay/getsignkey", paramXml)
	if err != nil {
		return
	}
	var keyResponse getSignKeyResponse
	if err = xml.Unmarshal(bytes, &keyResponse); err != nil {
		return
	}
	if keyResponse.ReturnCode == ResponseFail {
		err = errors.New(keyResponse.RetMsg)
		return
	}
	key = keyResponse.SandboxSignkey
	return
}
