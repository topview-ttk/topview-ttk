package login

import (
	"encoding/base64"
	"golang.org/x/crypto/argon2"
)

const (
	timeCost   uint32 = 3
	memoryCost uint32 = 32 * 1024
	threads    uint8  = 4
	keyLength  uint32 = 32
)

func EncryptPasswordWithSalt(password string, salt string) string {
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)

	hash := argon2.IDKey(passwordBytes, saltBytes, timeCost, memoryCost, threads, keyLength)
	hashBase64 := base64.RawStdEncoding.EncodeToString(hash)

	return hashBase64
}
