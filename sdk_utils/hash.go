package sdk_utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

// Md5 md5 hash string
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return hex.EncodeToString(m.Sum(nil))
}
