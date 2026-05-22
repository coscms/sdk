package sdk_options

type AppInfoGetter interface {
	GetAppSecret() string
	GetAppId() string
	GetApiEndpoint() string
}

type AppInfo struct {
	Secret      string
	AppId       string
	ApiEndpoint string
}

func (a AppInfo) GetAppSecret() string {
	return a.Secret
}

func (a AppInfo) GetAppId() string {
	return a.AppId
}

func (a AppInfo) GetApiEndpoint() string {
	return a.ApiEndpoint
}
