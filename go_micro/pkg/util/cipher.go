package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func Md5(b []byte) []byte {
	h := md5.New()
	h.Write(b)
	return h.Sum(nil)
}

func Sha256(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	return h.Sum(nil)
}

func Sha1(b []byte) []byte {
	h := sha1.New()
	h.Write(b)
	return h.Sum(nil)
}

func GetMd5Str(salt string) func(string) string {
	return func(s string) string {
		h := md5.New()
		io.WriteString(h, s)
		io.WriteString(h, salt)
		return hex.EncodeToString(h.Sum(nil))
	}
}

func GetSha1Str(salt string) func(string) string {
	return func(s string) string {
		h := sha1.New()
		io.WriteString(h, s)
		io.WriteString(h, salt)
		return hex.EncodeToString(h.Sum(nil))
	}
}

func GetSha256Str(salt string) func(string) string {
	return func(s string) string {
		h := sha256.New()
		io.WriteString(h, s)
		io.WriteString(h, salt)
		return hex.EncodeToString(h.Sum(nil))
	}
}

func GetHmacSha256Str(s string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}
