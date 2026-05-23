package sdk_options

import (
	"errors"
	"net/url"

	"github.com/coscms/sdk/sdk_utils"
)

// ToURLWithGenerator builds a signed URL with the given form data generator.
// This is the race-safe equivalent of SetGenerator + ToURL.
func (o *Options) ToURLWithGenerator(generator URLValuesGenerator, urlPath string, strength ...bool) (uri string, formData url.Values, err error) {
	if generator == nil {
		return "", nil, errors.New("generator is nil")
	}
	formData = generator.URLValues()
	appID := formData.Get(`appID`)
	if len(appID) == 0 {
		appID = o.GetAppID()
		formData.Set(`appID`, appID)
	} else {
		oldAppID := o.GetAppID()
		if oldAppID != appID {
			err = ErrAppIDConflict
			return
		}
	}
	appSecret := o.GetAppSecret()
	if len(strength) > 0 && strength[0] {
		appSecret = o.StrengthenSafeSecret(appSecret) // 加强防篡改安全性
	}
	formData = BuildURLValues(formData, appSecret, o.signaturer)
	uri = o.GetApiEndpoint() + urlPath
	return
}

// ToURL builds a signed URL using the pre-configured generator.
// Deprecated: Use ToURLWithGenerator to avoid data races on the shared generator state.
func (o *Options) ToURL(urlPath string, strength ...bool) (uri string, formData url.Values, err error) {
	return o.ToURLWithGenerator(o.generator, urlPath, strength...)
}

// StrengthenSafeSecret 强化的安全密钥(包含网络环境)
func (o *Options) StrengthenSafeSecret(secret string) string {
	return secret + `|` + sdk_utils.Md5(o.UserAgent) + `|` + o.ClientIP
}
