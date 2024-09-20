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
func EncPassword(password, salt string) string {
	return Md5Encode(password + salt)
}

// check password
func CheckPassword(password, salt, cipher string) bool {
	return EncPassword(password, salt) == cipher
}
