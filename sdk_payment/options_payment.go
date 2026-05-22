package sdk_payment

import (
	"net/url"
	"strings"

	"github.com/coscms/sdk/sdk_options"
)

// PaymentURL 付款接口GET提交地址
func (o *Options) PaymentURL(options *CheckoutOptions) (string, error) {
	urlString, formData, err := o.PaymentURLWithValues(options)
	if err != nil {
		return "", err
	}
	return urlString + `?` + formData.Encode(), nil
}

func (o *Options) PaystartURL(options *CheckoutOptions) (string, error) {
	urlString, formData, err := o.PaystartURLWithValues(options)
	if err != nil {
		return "", err
	}
	return urlString + `?` + formData.Encode(), nil
}

// PaymentProvierListURL 支付网关供应商列表
func (o *Options) PaymentProvierListURL(appID string, currencies ...string) (string, error) {
	urlValues := url.Values{}
	urlValues.Set(`appID`, appID)
	if len(currencies) > 0 {
		urlValues.Set(`currency`, strings.Join(currencies, `,`))
	}
	o.SetGenerator(sdk_options.DefaultURLValuesGenerator(urlValues))
	uri, formData, err := o.ToURL(`/open/v1/payment/providers`)
	if err != nil {
		return ``, err
	}
	return uri + `?` + formData.Encode(), nil
}

// PaymentURLWithValues /open/v1/payment/alipay
// 付款接口地址带表单数据
func (o *Options) PaymentURLWithValues(options *CheckoutOptions) (string, url.Values, error) {
	o.SetGenerator(options)
	return o.ToURL(`/open/v1/payment/` + options.Type)
}

// PaystartURLWithValues /open/v1/payment/start
// 付款选择方式页面网址
func (o *Options) PaystartURLWithValues(options *CheckoutOptions) (string, url.Values, error) {
	o.SetGenerator(options)
	return o.ToURL(`/open/v1/payment/start`)
}

// ClientPaymentQueryURLWithValues 构建客户端付款查询网址和参数值
func (o *Options) ClientPaymentQueryURLWithValues(appID string, outOrderNo string) (string, url.Values, error) {
	formData := url.Values{}
	formData.Set(`appID`, appID)
	formData.Set(`outOrderNo`, outOrderNo)
	o.SetGenerator(sdk_options.DefaultURLValuesGenerator(formData))
	return o.ToURL(`/open/v1/query/payment`)
}

func (o *Options) PaymentQueryURLWithValues(appID string, orderNo string, outOrderNo string) (string, url.Values, error) {
	formData := url.Values{}
	formData.Set(`appID`, appID)
	formData.Set(`orderNo`, orderNo)
	formData.Set(`outOrderNo`, outOrderNo)
	o.SetGenerator(sdk_options.DefaultURLValuesGenerator(formData))
	return o.ToURL(`/open/v1/query/payment`)
}

// RefundURL 退款接口GET提交地址
func (o *Options) RefundURL(options *RefundOptions) (string, error) {
	urlString, formData, err := o.RefundURLWithValues(options)
	if err != nil {
		return "", err
	}
	return urlString + `?` + formData.Encode(), nil
}

// RefundURLWithValues /open/v1/refund
// 退款接口地址带表单数据
func (o *Options) RefundURLWithValues(options *RefundOptions) (string, url.Values, error) {
	o.SetGenerator(options)
	return o.ToURL(`/open/v1/refund`)
}

// ClientRefundQueryURLWithValues 构建客户端退款查询网址和参数值
func (o *Options) ClientRefundQueryURLWithValues(appID string, outRefundNo string) (string, url.Values, error) {
	formData := url.Values{}
	formData.Set(`appID`, appID)
	formData.Set(`outRefundNo`, outRefundNo)
	o.SetGenerator(sdk_options.DefaultURLValuesGenerator(formData))
	return o.ToURL(`/open/v1/query/refund`)
}

func (o *Options) RefundQueryURLWithValues(appID string, refundNo string, outRefundNo string) (string, url.Values, error) {
	formData := url.Values{}
	formData.Set(`appID`, appID)
	formData.Set(`refundNo`, refundNo)
	formData.Set(`outRefundNo`, outRefundNo)
	o.SetGenerator(sdk_options.DefaultURLValuesGenerator(formData))
	return o.ToURL(`/open/v1/query/refund`)
}
