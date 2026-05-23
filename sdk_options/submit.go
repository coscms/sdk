package sdk_options

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/admpub/resty/v2"
	"github.com/coscms/sdk/sdk_utils"
	"github.com/webx-top/restyclient"
)

type Response struct {
	Code  int
	State string `json:",omitempty" xml:",omitempty"`
	Info  any
	URL   string `json:",omitempty" xml:",omitempty"`
	Zone  any    `json:",omitempty" xml:",omitempty"`
	Data  any    `json:",omitempty" xml:",omitempty"`
}

func (r Response) IsSuccess() bool {
	return r.Code == 1
}

func Submit(ctx context.Context, apiURL string, formData url.Values, method ...string) (*Response, error) {
	apiResp := &Response{}
	_, err := SubmitWithRecv(ctx, apiResp, apiURL, formData, method...)
	return apiResp, err
}

func SubmitWithRecv(ctx context.Context, recv interface{}, apiURL string, formData url.Values, method ...string) (*resty.Response, error) {
	request := restyclient.Retryable()
	request.SetContext(ctx)
	request.SetResult(recv)
	request.SetHeader(`Accept`, `application/json`)
	request.SetFormDataFromValues(formData)
	var resp *resty.Response
	var err error
	if len(method) > 0 && strings.EqualFold(method[0], `GET`) {
		resp, err = request.Get(apiURL)
	} else {
		resp, err = request.Post(apiURL)
	}
	if err != nil {
		if resp != nil {
			return nil, fmt.Errorf(`%w: %s: %s`, err, apiURL, sdk_utils.StripTags(resp.String()))
		}
		return nil, fmt.Errorf(`%w: %s`, err, apiURL)
	}
	if !resp.IsSuccess() {
		return resp, fmt.Errorf(`%s: %s: %s`, apiURL, resp.Status(), sdk_utils.StripTags(resp.String()))
	}
	return resp, err
}

func Submitx(ctx context.Context, apiURL string, formData url.Values, method ...string) (map[string]any, string, error) {
	apiResp := map[string]any{}
	var body string
	resp, err := SubmitWithRecv(ctx, &apiResp, apiURL, formData, method...)
	if resp != nil {
		body = sdk_utils.StripTags(resp.String())
	}
	return apiResp, body, err
}

func GetValueByKey(mp map[string]any, key string, fallbackKeys ...string) any {
	if len(mp) == 0 {
		return nil
	}
	val, ok := mp[key]
	if ok {
		return val
	}
	for _, fallback := range fallbackKeys {
		val, ok := mp[fallback]
		if ok {
			return val
		}
	}
	return nil
}
