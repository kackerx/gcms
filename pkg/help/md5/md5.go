package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}
