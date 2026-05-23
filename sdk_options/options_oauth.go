package sdk_options

import (
	"net/url"
)

// OauthProviderListURL 社区登录供应商列表
func (o *Options) OauthProviderListURL() (string, error) {
	urlValues := url.Values{}
	uri, formData, err := o.ToURLWithGenerator(DefaultURLValuesGenerator(urlValues), `/open/v1/oauth/providers`)
	if err != nil {
		return ``, err
	}
	return uri + `?` + formData.Encode(), nil
}

// OauthProvierListURL 保留旧名称以便向后兼容
// Deprecated: 请使用 OauthProviderListURL
func (o *Options) OauthProvierListURL() (string, error) {
	return o.OauthProviderListURL()
}
