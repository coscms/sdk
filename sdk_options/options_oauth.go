package sdk_options

import (
	"net/url"
)

// OAuthProviderListURL 社区登录供应商列表
func (o *Options) OAuthProviderListURL() (string, error) {
	urlValues := url.Values{}
	uri, formData, err := o.ToURLWithGenerator(DefaultURLValuesGenerator(urlValues), `/open/v1/oauth/providers`)
	if err != nil {
		return ``, err
	}
	return uri + `?` + formData.Encode(), nil
}
