package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// lowercase
func Md5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	tmp := h.Sum(nil)
	return hex.EncodeToString(tmp)
}

// uppercase
func MD5Encode(data string) string {
	return strings.ToUpper(Md5Encode(data))
}

// encrypt password
func MakePassword(password, salt string) string {
	return Md5Encode(password + salt)
}

// decrypt password
func CheckPassword(password, salt, hash string) bool {
	return MakePassword(password, salt) == hash
}
