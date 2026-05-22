package sdk_payment

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/coscms/sdk/sdk_options"
)

// NotifyOptions  通知参数
type NotifyOptions struct {
	// - App信息 -

	AppID string `json:"appID" xml:"appID"` // appID

	// - 订单信息 -

	// 订单号
	OrderNo string `json:"orderNo" xml:"orderNo"`
	// 业务方订单号
	OutOrderNo string  `json:"outOrderNo" xml:"outOrderNo"`
	Price      float64 `json:"price" xml:"price"`         // 售价
	RealPrice  float64 `json:"realPrice" xml:"realPrice"` // 实付金额
	PaidAt     uint    `json:"paidAt" xml:"paidAt"`       // 付款时间
	Extend     string  `json:"extend" xml:"extend"`       // 扩展信息

	// - 产品信息 -

	ProductID   string `json:"productID" xml:"productID"`     // 商品ID
	ProductType string `json:"productType" xml:"productType"` // 商品类型

	// 通知类型(payment-付款通知;refund-退款通知)
	Type    string `json:"type" xml:"type"`
	Subtype string `json:"subtype,omitempty" xml:"subtype,omitempty"` // 用于第四方支付时选择支付方式

	// - 退款信息 -

	// 退款单号
	RefundNo string `json:"refundNo,omitempty" xml:"refundNo,omitempty"`
	// 业务方退款单号
	OutRefundNo  string  `json:"outRefundNo,omitempty" xml:"outRefundNo,omitempty"`
	RefundAmount float64 `json:"refundAmount,omitempty" xml:"refundAmount,omitempty"` // 退款金额

	// - 通用：适用于付款和退款 -

	// 支付宝等支付网关平台的交易单号
	TransactionNo string `json:"transactionNo" xml:"transactionNo"`

	// - 状态 -

	Status      TradeStatus `json:"status,omitempty" xml:"status,omitempty"`           // 状态
	Description string      `json:"description,omitempty" xml:"description,omitempty"` // 说明

	Nonce string `json:"nonce,omitempty" xml:"nonce,omitempty"`
}

func (c *NotifyOptions) SetDefaults(get func(string) string) *NotifyOptions {
	if len(c.AppID) == 0 {
		c.AppID = get(`appId`)
	}
	if len(c.ProductID) == 0 {
		c.ProductID = get(`productId`)
	}
	return c
}

func (c *NotifyOptions) URLValues(apiKey string, signGenerators ...sdk_options.Signaturer) url.Values {
	params := url.Values{}
	params.Set(`appID`, c.AppID)
	params.Set(`orderNo`, c.OrderNo)
	params.Set(`outOrderNo`, c.OutOrderNo)
	params.Set(`price`, strconv.FormatFloat(c.Price, 'f', -1, 64))
	params.Set(`realPrice`, strconv.FormatFloat(c.RealPrice, 'f', -1, 64))
	params.Set(`type`, c.Type)
	params.Set(`paidAt`, strconv.FormatUint(uint64(c.PaidAt), 10))
	params.Set(`extend`, c.Extend)
	params.Set(`productID`, c.ProductID)
	params.Set(`productType`, c.ProductType)
	if len(c.Subtype) > 0 {
		params.Set(`subtype`, c.Subtype)
	}
	if len(c.Status) > 0 {
		params.Set(`status`, string(c.Status))
	}
	if len(c.Description) > 0 {
		params.Set(`description`, c.Description)
	}
	if len(c.OutRefundNo) > 0 {
		params.Set(`outRefundNo`, c.OutRefundNo)
		if len(c.RefundNo) > 0 {
			params.Set(`refundNo`, c.RefundNo)
		}
		params.Set(`refundAmount`, strconv.FormatFloat(c.RefundAmount, 'f', -1, 64))
	}
	if len(c.Nonce) > 0 {
		params.Set(`nonce`, c.Nonce)
	}

	var signGenerator sdk_options.Signaturer
	if len(signGenerators) > 0 {
		signGenerator = signGenerators[0]
	} else {
		signGenerator = sdk_options.GenSign
	}
	if signGenerator != nil {
		sign := signGenerator(params, apiKey)
		params.Set(`sign`, sign)
	}
	return params
}

func (c *NotifyOptions) Encode(apiKey string, signGenerators ...sdk_options.Signaturer) string {
	return c.URLValues(apiKey, signGenerators...).Encode()
}

func (c *NotifyOptions) IsSuccess() bool {
	return c.Status.IsSuccess()
}

func (c *NotifyOptions) IsFailure() bool {
	return c.Status.IsFailure()
}

func (c *NotifyOptions) IsCancelled() bool {
	return c.Status.IsCancelled()
}

func (c *NotifyOptions) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
