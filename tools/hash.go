package tools

import (
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

// WithKey set the key to sha256
func WithKey(str []byte) WithHashOption {
	return func(option *HashOption) {
		option.Key = str
	}
}

type WithHashOption func(option *HashOption)

func CheckSha1(expect string, actual []byte) bool {
	encoder := sha1.New()
	encoder.Write(actual)
	test := fmt.Sprintf("%x", encoder.Sum(nil))
	return expect == test
}
