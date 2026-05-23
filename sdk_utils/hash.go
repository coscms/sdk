package sdk_utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 computes the MD5 hex digest of the given string.
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}
