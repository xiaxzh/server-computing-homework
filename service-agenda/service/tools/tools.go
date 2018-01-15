package tools

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"

	uuid "github.com/satori/go.uuid"
)

// simple tools used in this program
// 1. encrypt password by MD5
// 2. creating UUID as id of user

// MD5Encryption to encrypt password by MD5
func MD5Encryption(text string) string {
	hash := md5.New()
	io.WriteString(hash, text)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

// GetKey creating UUID as key
func GetKey() string {
	return uuid.NewV4().String()
}

// GenenrateSessionID generate UID as session id
func GenenrateSessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}
