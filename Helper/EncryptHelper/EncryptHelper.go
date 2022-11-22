package EncryptHelper

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(data string, key string) []byte {
	var h = hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return h.Sum(nil)
}
