package wxpay

import (
	"net/http"
)

type Client struct {
	config       Config       // 配置信息
	serviceType  int          // 服务模式
	apiKey       string       // API Key
	certFilepath string       // 证书目录
	certClient   *http.Client // 带证书的http连接池
	isProd       bool         // 是否是生产环境
	isMch        bool         // 是否是特殊的商户接口(微信找零)
}

// 是否是服务商模式
func (c *Client) isFacilitator() bool {
	switch c.serviceType {
	case ServiceTypeFacilitatorDomestic, ServiceTypeFacilitatorAbroad, ServiceTypeBankServiceProvidor:
		return true
	default:
		return false
	}
}

// 拼接完整的URL
func (c *Client) url(relativePath string) string {
	if c.isProd {
		return baseUrl + relativePath
	} else {
		return baseUrlSandbox + relativePath
	}
}

// 初始化微信客户端
func NewClient(isProd bool, isMch bool, serviceType int, apiKey string, certFilepath string, config Config) (client *Client) {
	client = new(Client)
	client.config = config
	client.serviceType = serviceType
	client.apiKey = apiKey
	client.certFilepath = certFilepath
	client.isProd = isProd
	client.isMch = isMch
	return client
}
