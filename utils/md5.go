package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(string []byte) string {
	hash := md5.New()
	hash.Write(string)
	return hex.EncodeToString(hash.Sum(nil))
}