package wxpay

import (
	"encoding/xml"
)

// 统一下单
func (c *Client) UnifiedOrder(body UnifiedOrderBody) (wxRsp UnifiedOrderResponse, err error) {
	// 处理参数
	if body.SceneInfoModel != nil {
		body.SceneInfo = JsonString(*body.SceneInfoModel)
	}

	// 业务逻辑
	bytes, err := c.doWeChat("pay/unifiedorder", body)
	//fmt.Println(string(bytes))
	if err != nil {
		return
	}
	// 结果校验
	if err = c.doVerifySign(bytes, true); err != nil {
		return
	}
	// 解析返回值
	err = xml.Unmarshal(bytes, &wxRsp)
	return
}

// 统一下单的参数
type UnifiedOrderBody struct {
	SignType       string `json:"sign_type,omitempty"`   // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	DeviceInfo     string `json:"device_info,omitempty"` // (非必填) 终端设备号(门店号或收银设备ID)，注意：PC网页或JSAPI支付请传"WEB"
	Body           string `json:"body"`                  // 商品描述交易字段格式根据不同的应用场景建议按照以下格式上传： （1）PC网站——传入浏览器打开的网站主页title名-实际商品名称，例如：腾讯充值中心-QQ会员充值；（2） 公众号——传入公众号名称-实际商品名称，例如：腾讯形象店- image-QQ公仔；（3） H5——应用在浏览器网页上的场景，传入浏览器打开的移动网页的主页title名-实际商品名称，例如：腾讯充值中心-QQ会员充值；（4） 线下门店——门店品牌名-城市分店名-实际商品名称，例如： image形象店-深圳腾大- QQ公仔）（5） APP——需传入应用市场上的APP名字-实际商品名称，天天爱消除-游戏充值。
	Detail         string `json:"detail,omitempty"`      // (非必填) TODO 商品详细描述，对于使用单品优惠的商户，该字段必须按照规范上传，详见"单品优惠参数说明"
	Attach         string `json:"attach,omitempty"`      // (非必填) 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	OutTradeNo     string `json:"out_trade_no"`          // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*且在同一个商户号下唯一。详见商户订单号
	FeeType        string `json:"fee_type,omitempty"`    // (非必填) 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TotalFee       int    `json:"total_fee"`             // 订单总金额，单位为分，只能为整数，详见支付金额
	SpbillCreateIP string `json:"spbill_create_ip"`      // 支持IPV4和IPV6两种格式的IP地址。调用微信支付API的机器IP
	TimeStart      string `json:"time_start,omitempty"`  // (非必填) 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire     string `json:"time_expire,omitempty"` // (非必填) 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。订单失效时间是针对订单号而言的，由于在请求支付的时候有一个必传参数prepay_id只有两小时的有效期，所以在重入时间超过2小时的时候需要重新请求下单接口获取新的prepay_id。其他详见时间规则。建议：最短失效时间间隔大于1分钟
	GoodsTag       string `json:"goods_tag,omitempty"`   // (非必填) TODO 订单优惠标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	NotifyUrl      string `json:"notify_url"`            // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	TradeType      string `json:"trade_type"`            // JSAPI-JSAPI支付 NATIVE-Native支付 APP-APP支付 说明详见参数规定
	ProductId      string `json:"product_id,omitempty"`  // (非必填) trade_type=NATIVE时，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
	LimitPay       string `json:"limit_pay,omitempty"`   // (非必填) no_credit：指定不能使用信用卡支付
	OpenId         string `json:"openid,omitempty"`      // (非必填) trade_type=JSAPI，此参数必传，用户在主商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。下单前需要调用【网页授权获取用户信息】接口获取到用户的Openid。
	SubOpenId      string `json:"sub_openid,omitempty"`  // (非必填) trade_type=JSAPI，此参数必传，用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。下单前需要调用【网页授权获取用户信息】接口获取到用户的Openid。
	Receipt        string `json:"receipt,omitempty"`     // (非必填) Y，传入Y时，支付成功消息和支付详情页将出现开票入口。需要在微信支付商户平台或微信公众平台开通电子发票功能，传此字段才可生效
	SceneInfo      string `json:"scene_info,omitempty"`  // (非必填) 该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
	// 用于生成SceneInfo
	SceneInfoModel *SceneInfoModel `json:"-"`
}

// 统一下单的返回值
type UnifiedOrderResponse struct {
	ResponseModel
	// 当return_code为SUCCESS时
	ServiceResponseModel
	DeviceInfo string `xml:"device_info"` // 调用接口提交的终端设备号
	// 当return_code 和result_code都为SUCCESS时
	TradeType string `xml:"trade_type"` // JSAPI-公众号支付 NATIVE-Native支付 APP-APP支付 说明详见参数规定
	PrepayId  string `xml:"prepay_id"`  // 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	CodeUrl   string `xml:"code_url"`   // trade_type=NATIVE时有返回，此url用于生成支付二维码，然后提供给用户进行扫码支付。注意：code_url的值并非固定，使用时按照URL格式转成二维码即可
	MWebUrl   string `xml:"mweb_url"`   // mweb_url为拉起微信支付收银台的中间页面，可通过访问该url来拉起微信客户端，完成支付，mweb_url的有效期为5分钟。
}

//判断是否业务成功
func (resp UnifiedOrderResponse) ResultSuccess() bool {
	return resp.ReturnCode == ResponseSuccess && resp.ResultCode == ResponseSuccess
}
