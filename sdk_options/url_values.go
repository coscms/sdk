package sdk_options

import "net/url"

var DefaultURLValuesGenerator = func(values url.Values) URLValuesGenerator {
	return URLValues(values)
}

type URLValues url.Values

func (u URLValues) URLValues() url.Values {
	formData := url.Values(u)
	return formData
}
