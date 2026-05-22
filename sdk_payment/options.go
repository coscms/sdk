package sdk_payment

import (
	"net/url"

	"github.com/coscms/sdk/sdk_options"
)

// OptionsInterface 是 sdk_options.Options 中当前包所使用的函数的接口定义。
type OptionsInterface interface {
	SetGenerator(g sdk_options.URLValuesGenerator) *sdk_options.Options
	ToURL(urlPath string, strength ...bool) (string, url.Values, error)
}

func New(typ sdk_options.Type, appInfoGetter sdk_options.AppInfoGetter, opts ...func(*sdk_options.Options)) *Options {
	return &Options{
		OptionsInterface: sdk_options.New(typ, appInfoGetter, opts...),
	}
}

type Options struct {
	OptionsInterface
}
