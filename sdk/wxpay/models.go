package wxpay

import (
	"bytes"
	"fmt"
	"github.com/beevik/etree"
	"strconv"
)

// 微信支付的整体配置
type Config struct {
	AppId    string // 微信分配的公众账号ID
	SubAppId string // 微信分配的子商户公众账号ID
	MchId    string // 微信支付分配的商户号
	SubMchId string // 微信支付分配的子商户号，开发者模式下必填
}

// 场景信息模型
type SceneInfoModel struct {
	ID       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address"`   // 门店详细地址
}

// 返回结果的通信标识
type ResponseModel struct {
	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL 此字段是通信标识，非交易标识，交易是否成功需要查看result_code来判断
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因：签名失败/参数格式校验错误
	RetMsg     string `xml:"retmsg"`      // 沙盒时返回的错误信息
}

// 业务返回结果的错误信息
type ServiceResponseModel struct {
	AppId      string `xml:"appid"`        // 微信分配的公众账号ID
	MchId      string `xml:"mch_id"`       // 微信支付分配的商户号
	SubAppId   string `xml:"sub_appid"`    // (服务商模式) 微信分配的子商户公众账号ID
	SubMchId   string `xml:"sub_mch_id"`   // (服务商模式) 微信支付分配的子商户号
	NonceStr   string `xml:"nonce_str"`    // 随机字符串，不长于32位
	Sign       string `xml:"sign"`         // 签名，详见签名生成算法
	ResultCode string `xml:"result_code"`  // SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     // 详细参见第6节错误列表
	ErrCodeDes string `xml:"err_code_des"` // 错误返回的信息描述
}

// 特殊商户接口业务返回结果的错误信息
type MchServiceResponseModel struct {
	MchAppId   string `xml:"mch_appid"`    // 子商户公众账号ID
	MchId      string `xml:"mchid"`        // 子商户号
	NonceStr   string `xml:"nonce_str"`    // 随机字符串，不长于32位
	Sign       string `xml:"sign"`         // 签名，详见签名生成算法
	ResultCode string `xml:"result_code"`  // SUCCESS/FAIL
	ErrCode    string `xml:"err_code"`     // 详细参见第6节错误列表
	ErrCodeDes string `xml:"err_code_des"` // 错误返回的信息描述
}

// 微信通知的结果返回值
type NotifyResponseModel struct {
	ReturnCode string `xml:"return_code"` // SUCCESS/FAIL
	ReturnMsg  string `xml:"return_msg"`  // 返回信息，如非空，为错误原因，或OK
}

func (m *NotifyResponseModel) ToXmlString() string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("<xml>")
	buffer.WriteString(fmt.Sprintf("<return_code><![CDATA[%s]]></return_code>", m.ReturnCode))
	buffer.WriteString(fmt.Sprintf("<return_msg><![CDATA[%s]]></return_msg>", m.ReturnMsg))
	buffer.WriteString("</xml>")
	return buffer.String()
}

// 返回结果中的优惠券条目信息
type CouponResponseModel struct {
	CouponId   string // 代金券或立减优惠ID
	CouponType string // CASH-充值代金券 NO_CASH-非充值优惠券 开通免充值券功能，并且订单使用了优惠券后有返回
	CouponFee  int64  // 单个代金券或立减优惠支付金额
}

// 在XML节点树中，查找labels对应的
func NewCouponResponseModel(
	doc *etree.Element,
	idFormat string,
	typeFormat string,
	feeFormat string,
	numbers ...interface{},
) (m CouponResponseModel) {
	idName := fmt.Sprintf(idFormat, numbers...)
	typeName := fmt.Sprintf(typeFormat, numbers...)
	feeName := fmt.Sprintf(feeFormat, numbers...)
	m.CouponId = doc.SelectElement(idName).Text()
	m.CouponType = doc.SelectElement(typeName).Text()
	m.CouponFee, _ = strconv.ParseInt(doc.SelectElement(feeName).Text(), 10, 64)
	return
}
