package sdk_options

import "net/url"

var DefaultURLValuesGenerator = func(values url.Values) URLValuesGenerator {
	return URLValues(values)
}

type URLValues url.Values

func (u URLValues) URLValues(apiKey string, signGenerators ...Signaturer) url.Values {
	formData := url.Values(u)
	var signGenerator Signaturer
	if len(signGenerators) > 0 {
		signGenerator = signGenerators[0]
	} else {
		signGenerator = GenSign
	}
	if signGenerator != nil {
		sign := signGenerator(formData, apiKey)
		formData.Set(`sign`, sign)
	}
	return formData
}
