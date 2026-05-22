package sdk_options

import (
	"net/url"
	"strconv"
	"time"

	"github.com/coscms/stdauth"
)

// BuildURLValues 构建API参数值
func BuildURLValues(values url.Values, secret string) url.Values {
	if values == nil {
		values = url.Values{}
	}
	//values.Set(`appID`, appID)
	values.Set(`timestamp`, strconv.FormatInt(time.Now().Unix(), 10))
	values.Del(`sign`)
	sign := stdauth.MakeSign(values, secret)
	values.Set(`sign`, sign)
	return values
}
