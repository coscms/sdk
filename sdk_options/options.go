package sdk_options

import (
	"net/url"
	"strings"
)

func OptionType(typ Type) func(*Options) {
	return func(o *Options) {
		o.Type = typ
	}
}

func OptionUserAgent(userAgent string) func(*Options) {
	return func(o *Options) {
		o.UserAgent = userAgent
	}
}

func OptionClientIP(clientIP string) func(*Options) {
	return func(o *Options) {
		o.ClientIP = clientIP
	}
}

func OptionAppInfoGetter(appInfoGetter AppInfoGetter) func(*Options) {
	return func(o *Options) {
		o.appInfoGetter = appInfoGetter
	}
}

func OptionSignaturer(signaturer Signaturer) func(*Options) {
	return func(o *Options) {
		o.signaturer = signaturer
	}
}

func New(typ Type, appInfoGetter AppInfoGetter, opts ...func(*Options)) *Options {
	o := &Options{
		Type:          typ,
		appInfoGetter: appInfoGetter,
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

type Signaturer func(url.Values, string) string

type Options struct {
	generator     URLValuesGenerator
	signaturer    Signaturer
	appInfoGetter AppInfoGetter
	Type          Type
	UserAgent     string
	ClientIP      string
}

func (o *Options) SetGenerator(g URLValuesGenerator) *Options {
	o.generator = g
	return o
}

func (o *Options) SetSignaturer(fn Signaturer) *Options {
	o.signaturer = fn
	return o
}

func (o *Options) SetAppInfoGetter(appInfoGetter AppInfoGetter) *Options {
	o.appInfoGetter = appInfoGetter
	return o
}

func (o *Options) GetAppInfoGetter() AppInfoGetter {
	return o.appInfoGetter
}

func (o *Options) GetAppID() string {
	if o.appInfoGetter != nil {
		return o.appInfoGetter.GetAppId()
	}
	return ``
}

func (o *Options) GetAppSecret() string {
	if o.appInfoGetter != nil {
		return o.appInfoGetter.GetAppSecret()
	}
	return ``
}

func (o *Options) GetApiEndpoint() string {
	if o.appInfoGetter != nil {
		return strings.TrimSuffix(o.appInfoGetter.GetApiEndpoint(), `/`)
	}
	return ``
}
