package sdk_options

import (
	"net/url"

	"github.com/coscms/stdauth"
)

// SignString generates a signature string from raw content and apiKey.
func SignString(raw string, apiKey string) string {
	return stdauth.SignString(raw, apiKey)
}

// CheckSign 检查签名是否匹配
func CheckSign(raw string, sign string, apiKey string) error {
	if SignString(raw, apiKey) != sign {
		return ErrInvalidSign
	}
	return nil
}

// GenSign 根据url.Values类型值生成签名
func GenSign(formData url.Values, apiKey string) string {
	formData.Del(`sign`)
	return stdauth.MakeSign(formData, apiKey)
}
