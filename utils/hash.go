package utils

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

// Md5 生成 md5 字符串
func Md5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

func Hash(hash hash.Hash, data []byte) string {
	hash.Write(data)
	return hex.EncodeToString(hash.Sum(nil))
}
