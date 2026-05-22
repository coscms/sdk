package sdk_options

type AppInfoGetter interface {
	GetAppSecret() string
	GetAppID() string
	GetApiEndpoint() string
}

type AppInfo struct {
	Secret      string
	AppID       string
	ApiEndpoint string
}

func (a AppInfo) GetAppSecret() string {
	return a.Secret
}

func (a AppInfo) GetAppID() string {
	return a.AppID
}

func (a AppInfo) GetApiEndpoint() string {
	return a.ApiEndpoint
}
