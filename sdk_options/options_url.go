package sdk_options

import (
	"net/url"

	"github.com/coscms/sdk/sdk_utils"
)

func (o *Options) ToURL(urlPath string, strength ...bool) (uri string, formData url.Values, err error) {
	formData = o.generator.URLValues(o.GetAppSecret(), o.signaturer)
	appID := formData.Get(`appID`)
	if len(appID) == 0 {
		appID = o.GetAppId()
		formData.Set(`appID`, appID)
	} else {
		oldAppID := o.GetAppId()
		if oldAppID != appID {
			err = ErrAppIDConflict
			return
		}
	}
	appSecret := o.GetAppSecret()
	if len(strength) > 0 && strength[0] {
		appSecret = o.StrengthenSafeSecret(appSecret) // 加强防篡改安全性
	}
	formData = BuildURLValues(formData, appSecret)
	uri = o.GetApiEndpoint() + urlPath
	return
}

// StrengthenSafeSecret 强化的安全密钥(包含网络环境)
func (o *Options) StrengthenSafeSecret(secret string) string {
	return secret + `|` + sdk_utils.Md5(o.UserAgent) + `|` + o.ClientIP
}

// SecretSafeBuilder 密钥强化
func (o *Options) SecretSafeBuilder(secret string) string {
	return secret + `|` + sdk_utils.Md5(o.UserAgent)
}
