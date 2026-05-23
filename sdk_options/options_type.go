package sdk_options

import "net/url"

// URLValuesGenerator is the interface that wraps the URLValues method.
// URLValues generates form values for an API request.
type URLValuesGenerator interface {
	URLValues() url.Values
}

// Type represents the type of SDK options.
type Type string

const (
	// TypeOauth 社区登录类型
	TypeOauth Type = "oauth"
	// TypePayment 支付类型
	TypePayment Type = "payment"
)
