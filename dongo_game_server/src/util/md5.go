package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 前后端一致的算法即可
func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
