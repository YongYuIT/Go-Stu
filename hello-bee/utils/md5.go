package utils

import (
	"crypto/md5"
	"encoding/hex"
	uuid "github.com/satori/go.uuid"
)

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Uuid() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		return ""
	}
	return uuid.String()
}
