package sdk_options

import "net/url"

// DefaultURLValuesGenerator wraps a url.Values as a URLValuesGenerator.
var DefaultURLValuesGenerator = func(values url.Values) URLValuesGenerator {
	return URLValues(values)
}

// URLValues is a url.Values wrapper that implements URLValuesGenerator.
type URLValues url.Values

// URLValues implements the URLValuesGenerator interface.
func (u URLValues) URLValues() url.Values {
	formData := url.Values(u)
	return formData
}
