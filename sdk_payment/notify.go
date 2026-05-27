package sdk_payment

import (
	"encoding/json"
	"errors"
	"net/url"
	"strconv"
)

// NotifyOptions 通知参数
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
	Type string `json:"type" xml:"type"`

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

// SetDefaults fills empty fields with values from the given config function.
func (n *NotifyOptions) SetDefaults(get func(string) string) *NotifyOptions {
	if len(n.AppID) == 0 {
		n.AppID = get(`appId`)
	}
	if len(n.ProductID) == 0 {
		n.ProductID = get(`productId`)
	}
	return n
}

// URLValues serializes the notify options to url.Values.
func (n *NotifyOptions) URLValues() url.Values {
	params := url.Values{}
	params.Set(`appID`, n.AppID)
	params.Set(`orderNo`, n.OrderNo)
	params.Set(`outOrderNo`, n.OutOrderNo)
	params.Set(`price`, strconv.FormatFloat(n.Price, 'f', -1, 64))
	params.Set(`realPrice`, strconv.FormatFloat(n.RealPrice, 'f', -1, 64))
	params.Set(`type`, n.Type)
	params.Set(`paidAt`, strconv.FormatUint(uint64(n.PaidAt), 10))
	params.Set(`extend`, n.Extend)
	params.Set(`productID`, n.ProductID)
	params.Set(`productType`, n.ProductType)
	if len(n.Status) > 0 {
		params.Set(`status`, string(n.Status))
	}
	if len(n.Description) > 0 {
		params.Set(`description`, n.Description)
	}
	if len(n.OutRefundNo) > 0 {
		params.Set(`outRefundNo`, n.OutRefundNo)
		if len(n.RefundNo) > 0 {
			params.Set(`refundNo`, n.RefundNo)
		}
		params.Set(`refundAmount`, strconv.FormatFloat(n.RefundAmount, 'f', -1, 64))
	}
	if len(n.Nonce) > 0 {
		params.Set(`nonce`, n.Nonce)
	}
	return params
}

// IsSuccess returns true if the notify status is success.
func (n *NotifyOptions) IsSuccess() bool {
	return n.Status.IsSuccess()
}

// IsFailure returns true if the notify status is failure.
func (n *NotifyOptions) IsFailure() bool {
	return n.Status.IsFailure()
}

// IsCancelled returns true if the notify status is cancelled.
func (n *NotifyOptions) IsCancelled() bool {
	return n.Status.IsCancelled()
}

// String returns the JSON representation of NotifyOptions.
func (n *NotifyOptions) String() string {
	b, _ := json.Marshal(n)
	return string(b)
}

func (n *NotifyOptions) Validate() error {
	if len(n.AppID) == 0 {
		return errors.New(`appID is required`)
	}
	if len(n.Status) == 0 {
		return errors.New(`status is required`)
	}
	switch n.Type {
	case `payment`:
		if len(n.OutOrderNo) == 0 {
			return errors.New(`outOrderNo is required`)
		}
		if n.RealPrice <= 0 {
			return errors.New(`realPrice is required`)
		}

	case `refund`:
		if len(n.OutRefundNo) == 0 {
			return errors.New(`outRefundNo is required`)
		}
		if n.RefundAmount <= 0 {
			return errors.New(`refundAmount is required`)
		}

	default:
	}
	return nil
}
