package sdk_options

import "errors"

var (
	ErrAppIDConflict = errors.New(`AppID前后不一致`)
	ErrInvalidSign   = errors.New(`签名无效`)
)
