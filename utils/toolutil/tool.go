package toolutil

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GenerateToken(ip, userAgent string) string {
	return Md5(ip+":"+userAgent) + Md5(time.Now().String())
}
