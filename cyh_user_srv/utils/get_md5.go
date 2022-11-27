package utils

import (
	"crypto/md5"
	"encoding/hex"
)

var salt = "chenyuhan.*saxc"

func Md5Pwd(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd + salt))
	return hex.EncodeToString(h.Sum(nil))
}
