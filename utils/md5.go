package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(string []byte, b ...byte) string {
	h := md5.New()
	h.Write(string)
	return hex.EncodeToString(h.Sum(b))
}
