package sdk_options

import (
	"net/url"
	"strconv"
	"time"
)

// BuildURLValues 构建API参数值
func BuildURLValues(values url.Values, secret string, signaturer Signaturer) url.Values {
	if values == nil {
		values = url.Values{}
	}
	//values.Set(`appID`, appID)
	values.Set(`timestamp`, strconv.FormatInt(time.Now().Unix(), 10))
	values.Del(`sign`)
	if signaturer == nil {
		signaturer = GenSign
	}
	sign := signaturer(values, secret)
	values.Set(`sign`, sign)
	return values
}
