package tools

import (
	"crypto"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type HashOption struct {
	Key []byte
}

// Hash make hash use sha256
func Hash(str string, options ...WithHashOption) (string, error) {
	var option HashOption
	for _, o := range options {
		o(&option)
	}
	cryptor := sha256.New()
	_, err := cryptor.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(cryptor.Sum(option.Key)), nil
}

//CheckHash check if hash and text string are equaled.
func CheckHash(hashStr string, text string, options ...WithHashOption) bool {
	hash, err := Hash(text, options...)
	if err != nil {
		return false
	}
	return hash == hashStr
}

// WithKey set the key to sha256
func WithKey(str []byte) WithHashOption {
	return func(option *HashOption) {
		option.Key = str
	}
}

type WithHashOption func(option *HashOption)

//Sha1 sha1算法
func Sha1(str string) string {
	hash := crypto.SHA1.New()
	hash.Write([]byte(str))
	return hex.EncodeToString(hash.Sum(nil))
}

func CheckSha1(expect string, actual []byte) bool {
	encoder := sha1.New()
	encoder.Write(actual)
	test := fmt.Sprintf("%x", encoder.Sum(nil))
	return expect == test
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
