package iutil

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5Encrypt(noSignStr string) string {
	h := md5.New()
	h.Write([]byte(noSignStr))
	// 生成的签名
	return hex.EncodeToString(h.Sum(nil))
}
