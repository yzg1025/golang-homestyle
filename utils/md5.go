package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func MD5(string []byte) string {
	hash := md5.New()
	hash.Write(string)
	return hex.EncodeToString(hash.Sum(nil))
}

func Generate() string {
	uid := uuid.NewV4()
	return fmt.Sprintf("api_%v", uid)
}
