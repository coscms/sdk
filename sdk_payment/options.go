package sdk_payment

import (
	"github.com/coscms/sdk/sdk_options"
)

func New(typ sdk_options.Type, appInfoGetter sdk_options.AppInfoGetter, opts ...func(*sdk_options.Options)) *Options {
	return &Options{
		Options: sdk_options.New(typ, appInfoGetter, opts...),
	}
}

type Options struct {
	*sdk_options.Options
}
