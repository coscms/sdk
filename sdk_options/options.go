package sdk_options

import (
	"net/url"
	"strings"
)

// OptionType sets the option type for Options.
func OptionType(typ Type) func(*Options) {
	return func(o *Options) {
		o.Type = typ
	}
}

// OptionUserAgent sets the user agent for Options.
func OptionUserAgent(userAgent string) func(*Options) {
	return func(o *Options) {
		o.UserAgent = userAgent
	}
}

// OptionClientIP sets the client IP for Options.
func OptionClientIP(clientIP string) func(*Options) {
	return func(o *Options) {
		o.ClientIP = clientIP
	}
}

// OptionAppInfoGetter sets the app info getter for Options.
func OptionAppInfoGetter(appInfoGetter AppInfoGetter) func(*Options) {
	return func(o *Options) {
		o.appInfoGetter = appInfoGetter
	}
}

// OptionSigner sets a custom signature function.
func OptionSigner(signer Signer) func(*Options) {
	return func(o *Options) {
		o.signaturer = signer
	}
}

// Deprecated: Use OptionSigner instead.
func OptionSignaturer(signaturer Signaturer) func(*Options) {
	return OptionSigner(signaturer)
}

// New creates a new Options with the given type and app info.
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

// Signer generates a signature from form values and a secret key.
type Signer func(url.Values, string) string

// Deprecated: Use Signer instead.
type Signaturer = Signer

// Options holds the configuration for signing and submitting API requests.
type Options struct {
	generator     URLValuesGenerator
	signaturer    Signaturer
	appInfoGetter AppInfoGetter
	Type          Type
	UserAgent     string
	ClientIP      string
}

// SetGenerator sets the URL values generator.
func (o *Options) SetGenerator(g URLValuesGenerator) *Options {
	o.generator = g
	return o
}

// SetSigner sets a custom signature function.
func (o *Options) SetSigner(fn Signer) *Options {
	o.signaturer = fn
	return o
}

// Deprecated: Use SetSigner instead.
func (o *Options) SetSignaturer(fn Signaturer) *Options {
	return o.SetSigner(fn)
}

// SetAppInfoGetter sets the app info getter.
func (o *Options) SetAppInfoGetter(appInfoGetter AppInfoGetter) *Options {
	o.appInfoGetter = appInfoGetter
	return o
}

// GetAppInfoGetter returns the app info getter.
func (o *Options) GetAppInfoGetter() AppInfoGetter {
	return o.appInfoGetter
}

// GetAppID returns the app ID from the app info getter.
func (o *Options) GetAppID() string {
	if o.appInfoGetter != nil {
		return o.appInfoGetter.GetAppID()
	}
	return ``
}

// GetAppSecret returns the app secret from the app info getter.
func (o *Options) GetAppSecret() string {
	if o.appInfoGetter != nil {
		return o.appInfoGetter.GetAppSecret()
	}
	return ``
}

// GetApiEndpoint returns the API endpoint with trailing slash removed.
func (o *Options) GetApiEndpoint() string {
	if o.appInfoGetter != nil {
		return strings.TrimSuffix(o.appInfoGetter.GetApiEndpoint(), `/`)
	}
	return ``
}
