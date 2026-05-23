package sdk_options

import "errors"

var (
	// ErrAppIDConflict is returned when the appID in the form data conflicts with the configured appID.
	ErrAppIDConflict = errors.New(`AppID前后不一致`)
	// ErrInvalidSign is returned when the signature verification fails.
	ErrInvalidSign = errors.New(`签名无效`)
)
