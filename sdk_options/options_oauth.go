package sdk_options

import (
	"net/url"
)

// OauthProvierListURL 社区登录供应商列表
func (o *Options) OauthProvierListURL() (string, error) {
	urlValues := url.Values{}
	o.SetGenerator(DefaultURLValuesGenerator(urlValues))
	uri, formData, err := o.ToURL(`/open/v1/oauth/providers`)
	if err != nil {
		return ``, err
	}
	return uri + `?` + formData.Encode(), nil
}
