package sdk_payment

import (
	"encoding/json"
	"net/url"
	"strconv"

	"github.com/coscms/sdk/sdk_options"
)

// CheckoutOptions  付款参数
type CheckoutOptions struct {
	// - App信息 -

	AppID  string `json:"appID" xml:"appID" valid:"required" validate:"required"` // appID
	AppUID string `json:"appUID" xml:"appUID"`                                    // 你的用户ID

	// - 订单信息 -

	// 币种 支持的值：https://github.com/webx-top/payment/blob/master/config/constants.go
	Currency   string  `json:"currency" xml:"currency"`
	Price      float64 `json:"price" xml:"price" valid:"required;min(0.01)" validate:"required"` // 价格
	OutOrderNo string  `json:"outOrderNo" xml:"outOrderNo" valid:"required" validate:"required"` // 你的订单号
	Subject    string  `json:"subject" xml:"subject" valid:"required" validate:"required"`       // 订单主题(一般为商品名)
	ExpiresTs  uint64  `json:"expiresTs" xml:"expiresTs"`                                        // 过期时间戳(秒)
	Extend     string  `json:"extend" xml:"extend"`                                              // 扩展信息

	// - 网址信息 -

	NotifyURL string `json:"notifyURL" xml:"notifyURL"` // 支付回调
	ReturnURL string `json:"returnURL" xml:"returnURL"` // 支付成功后返回地址
	CancelURL string `json:"cancelURL" xml:"cancelURL"` // 支付放弃后返回地址

	// - 产品信息 -

	ProductID        string `json:"productID" xml:"productID"`               // 商品ID
	ProductType      string `json:"productType" xml:"productType"`           // 商品类型(自己定义)
	IsVirtualProduct bool   `json:"isVirtualProduct" xml:"isVirtualProduct"` // 是否是虚拟商品

	// 付款方式(alipay,wechat)
	Type    string `json:"type" xml:"type" valid:"required" validate:"required"`
	Subtype string `json:"subtype,omitempty" xml:"subtype,omitempty"` // 用于第四方支付时选择支付方式
	// 设备 支持的值：https://github.com/webx-top/payment/blob/master/config/constants.go
	Device string `json:"device" xml:"device"`
	// 客户IP
	CustomerIP string `json:"customerIP" xml:"customerIP"`

	Nonce string `json:"nonce,omitempty" xml:"nonce,omitempty"`
}

func (c *CheckoutOptions) SetDefaults(get func(string) string) *CheckoutOptions {
	if len(c.AppID) == 0 {
		c.AppID = get(`appId`)
	}
	if len(c.AppUID) == 0 {
		c.AppUID = get(`appUid`)
	}
	if len(c.NotifyURL) == 0 {
		c.NotifyURL = get(`notifyUrl`)
	}
	if len(c.ReturnURL) == 0 {
		c.ReturnURL = get(`returnUrl`)
	}
	if len(c.ProductID) == 0 {
		c.ProductID = get(`productId`)
	}
	if len(c.CustomerIP) == 0 {
		c.CustomerIP = get(`customerIp`)
	}
	return c
}

func (c *CheckoutOptions) URLValues(apiKey string, signGenerators ...sdk_options.Signaturer) url.Values {
	formData := url.Values{}
	formData.Set(`appID`, c.AppID)
	formData.Set(`appUID`, c.AppUID)
	formData.Set(`productID`, c.ProductID)
	formData.Set(`productType`, c.ProductType)
	if c.IsVirtualProduct {
		formData.Set(`isVirtualProduct`, `1`)
	} else {
		formData.Set(`isVirtualProduct`, `0`)
	}
	formData.Set(`currency`, c.Currency)
	formData.Set(`price`, strconv.FormatFloat(c.Price, 'f', -1, 64))
	formData.Set(`outOrderNo`, c.OutOrderNo)
	formData.Set(`extend`, c.Extend)
	formData.Set(`notifyURL`, c.NotifyURL)
	formData.Set(`returnURL`, c.ReturnURL)
	formData.Set(`cancelURL`, c.CancelURL)
	formData.Set(`subject`, c.Subject)
	formData.Set(`expiresTs`, strconv.FormatUint(c.ExpiresTs, 10))
	formData.Set(`type`, c.Type)
	if len(c.Subtype) > 0 {
		formData.Set(`subtype`, c.Subtype)
	}
	formData.Set(`device`, c.Device)
	formData.Set(`customerIP`, c.CustomerIP)
	if len(c.Nonce) > 0 {
		formData.Set(`nonce`, c.Nonce)
	}

	var signGenerator func(url.Values, string) string
	if len(signGenerators) > 0 {
		signGenerator = signGenerators[0]
	} else {
		signGenerator = sdk_options.GenSign
	}
	if signGenerator != nil {
		sign := signGenerator(formData, apiKey)
		formData.Set(`sign`, sign)
	}
	return formData
}

func (c *CheckoutOptions) Encode(apiKey string, signGenerators ...sdk_options.Signaturer) string {
	return c.URLValues(apiKey, signGenerators...).Encode()
}

func (c *CheckoutOptions) String() string {
	b, _ := json.Marshal(c)
	return string(b)
}
