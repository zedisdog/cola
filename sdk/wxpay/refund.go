package wxpay

import (
	"encoding/xml"
	"errors"
	"github.com/beevik/etree"
)

// 申请退款
func (c *Client) Refund(body RefundBody) (wxRsp RefundResponse, err error) {
	// 业务逻辑
	bytes, err := c.doWeChatWithCert("secapi/pay/refund", body)
	if err != nil {
		return
	}
	// 结果校验
	if err = c.doVerifySign(bytes, true); err != nil {
		return
	}
	// 解析返回值
	err = c.refundParseResponse(bytes, &wxRsp)
	return
}

// 申请退款的参数
type RefundBody struct {
	TransactionId string `json:"transaction_id"`            // 微信支付订单号
	OutTradeNo    string `json:"out_trade_no"`              // 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。
	OutRefundNo   string `json:"out_refund_no"`             // 商户系统内部的退款单号，商户系统内部唯一，只能是数字、大小写字母_-|*@ ，同一退款单号多次请求只退一笔。
	TotalFee      int    `json:"total_fee"`                 // 订单总金额，单位为分，只能为整数，详见支付金额
	RefundFee     int    `json:"refund_fee"`                // 退款总金额，单位为分，只能为整数，可部分退款。详见支付金额
	RefundFeeType string `json:"refund_fee_type,omitempty"` // 退款货币类型，需与支付一致，或者不填。符合ISO 4217标准的三位字母代码，默认人民币：CNY
	RefundDesc    string `json:"refund_desc,omitempty"`     // 现退款原因
	RefundAccount string `json:"refund_account,omitempty"`  // 退款资金来源(见constatnt定义)
	NotifyUrl     string `json:"notify_url,omitempty"`      // 异步接收微信支付退款结果通知的回调地址
}

// 申请退款的返回值
type RefundResponse struct {
	ResponseModel
	ServiceResponseModel
	TransactionId       string `xml:"transaction_id"`
	OutTradeNo          string `xml:"out_trade_no"`
	OutRefundNo         string `xml:"out_refund_no"`
	RefundId            string `xml:"refund_id"`
	RefundFee           int    `xml:"refund_fee"`
	SettlementRefundFee int    `xml:"settlement_refund_fee"`
	TotalFee            int    `xml:"total_fee"`
	SettlementTotalFee  int    `xml:"settlement_total_fee"`
	FeeType             string `xml:"fee_type"`
	CashFee             int    `xml:"cash_fee"`
	CashRefundFee       int    `xml:"cash_refund_fee"`
	CouponRefundFee     int    `xml:"coupon_refund_fee"`
	CouponRefundCount   int    `xml:"coupon_refund_count"`
	// 使用coupon_refund_count的序号生成的优惠券项
	RefundCoupons []CouponResponseModel `xml:"-"`
}

// 申请退款-解析XML返回值
func (c *Client) refundParseResponse(xmlStr []byte, rsp *RefundResponse) (err error) {
	// 常规解析
	if err = xml.Unmarshal(xmlStr, rsp); err != nil {
		return
	}
	// 解析CouponRefundCount的对应项
	if rsp.CouponRefundCount > 0 {
		doc := etree.NewDocument()
		if err = doc.ReadFromBytes(xmlStr); err != nil {
			return
		}
		root := doc.SelectElement("xml")
		if root == nil {
			err = errors.New("xml格式错误")
			return
		}
		for i := 0; i < rsp.CouponRefundCount; i++ {
			m := NewCouponResponseModel(root, "coupon_refund_id_%d", "coupon_type_%d", "coupon_refund_fee_%d", i)
			rsp.RefundCoupons = append(rsp.RefundCoupons, m)
		}
	}
	return
}
