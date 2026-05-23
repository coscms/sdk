package sdk_options

// AppInfoGetter provides application credentials for API authentication.
type AppInfoGetter interface {
	GetAppSecret() string
	GetAppID() string
	GetApiEndpoint() string
}

// AppInfo is a concrete implementation of AppInfoGetter.
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
