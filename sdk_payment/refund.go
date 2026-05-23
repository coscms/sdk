package sdk_payment

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// RefundOptions 退款参数
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

// SetDefaults fills empty fields with values from the given config function.
func (r *RefundOptions) SetDefaults(get func(string) string) *RefundOptions {
	if len(r.AppID) == 0 {
		r.AppID = get(`appId`)
	}
	if len(r.NotifyURL) == 0 {
		r.NotifyURL = get(`notifyUrl`)
	}
	return r
}

// URLValues serializes the refund options to url.Values.
func (r *RefundOptions) URLValues() url.Values {
	formData := url.Values{}
	formData.Set(`appID`, r.AppID)
	formData.Set(`orderNo`, r.OrderNo)
	formData.Set(`outOrderNo`, r.OutOrderNo)
	formData.Set(`refundAmount`, strconv.FormatFloat(r.RefundAmount, 'f', -1, 64))
	formData.Set(`outRefundNo`, r.OutRefundNo)
	formData.Set(`refundReason`, r.RefundReason)
	formData.Set(`notifyURL`, r.NotifyURL)
	if r.AlwaysSave != nil {
		formData.Set(`alwaysSave`, strconv.FormatBool(*r.AlwaysSave))
	}
	if len(r.Nonce) > 0 {
		formData.Set(`nonce`, r.Nonce)
	}
	return formData
}

// String returns the JSON representation of RefundOptions.
func (r *RefundOptions) String() string {
	b, _ := json.Marshal(r)
	return string(b)
}
