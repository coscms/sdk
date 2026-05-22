package sdk_options

import "net/url"

type URLValuesGenerator interface {
	URLValues() url.Values
}

type Type string

const (
	// TypeOauth 社区登录类型
	TypeOauth Type = "oauth"
	// TypePayment 支付类型
	TypePayment Type = "payment"
)
