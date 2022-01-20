package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Md5 生成 md5 字符串
func Md5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
