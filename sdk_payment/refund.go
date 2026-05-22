package sdk_payment

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// RefundOptions  退款参数
type RefundOptions struct {
	// - App信息 -

	AppID string `json:"appID" xml:"appID" valid:"required" validate:"required"` // appID

	// - 订单信息 -

	OrderNo    string `json:"orderNo" xml:"orderNo"`                                            // 平台订单号
	OutOrderNo string `json:"outOrderNo" xml:"outOrderNo" valid:"required" validate:"required"` // 你的订单号

	// - 退款信息 -

	RefundAmount float64 `json:"refundAmount" xml:"refundAmount" valid:"required;min(0.01)" validate:"required"` // 退款金额
	OutRefundNo  string  `json:"outRefundNo" xml:"outRefundNo"`                                                  // 你的退款单号
	RefundReason string  `json:"refundReason" xml:"refundReason"`                                                // 退款原因
	NotifyURL    string  `json:"notifyURL" xml:"notifyURL"`                                                      // 通知接口网址

	AlwaysSave *bool `json:"alwaysSave,omitempty" xml:"alwaysSave,omitempty"` // 是否在支付网关不支持退款时依然保存退款单

	Nonce string `json:"nonce,omitempty" xml:"nonce,omitempty"`
}

func (c *RefundOptions) SetDefaults(get func(string) string) *RefundOptions {
	if len(c.AppID) == 0 {
		c.AppID = get(`appId`)
	}
	if len(c.NotifyURL) == 0 {
		c.NotifyURL = get(`notifyUrl`)
	}
	return c
}

func (c *RefundOptions) URLValues() url.Values {
	formData := url.Values{}
	formData.Set(`appID`, c.AppID)
	formData.Set(`orderNo`, c.OrderNo)
	formData.Set(`outOrderNo`, c.OutOrderNo)
	formData.Set(`refundAmount`, strconv.FormatFloat(c.RefundAmount, 'f', -1, 64))
	formData.Set(`outRefundNo`, c.OutRefundNo)
	formData.Set(`refundReason`, c.RefundReason)
	formData.Set(`notifyURL`, c.NotifyURL)
	if c.AlwaysSave != nil {
		formData.Set(`alwaysSave`, strconv.FormatBool(*c.AlwaysSave))
	}
	if len(c.Nonce) > 0 {
		formData.Set(`nonce`, c.Nonce)
	}
	return formData
}

func (c *RefundOptions) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
