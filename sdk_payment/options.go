package sdk_payment

import (
	"net/url"

	"github.com/coscms/sdk/sdk_options"
)

// OptionsInterface defines the subset of sdk_options.Options methods used by this package.
type OptionsInterface interface {
	SetGenerator(g sdk_options.URLValuesGenerator) *sdk_options.Options
	ToURL(urlPath string, strength ...bool) (string, url.Values, error)
	ToURLWithGenerator(generator sdk_options.URLValuesGenerator, urlPath string, strength ...bool) (string, url.Values, error)
}

// New creates a new payment Options with the given type and app info.
func New(typ sdk_options.Type, appInfoGetter sdk_options.AppInfoGetter, opts ...func(*sdk_options.Options)) *Options {
	return &Options{
		OptionsInterface: sdk_options.New(typ, appInfoGetter, opts...),
	}
}

// Options wraps sdk_options.Options to provide payment-specific URL building methods.
type Options struct {
	OptionsInterface
}
