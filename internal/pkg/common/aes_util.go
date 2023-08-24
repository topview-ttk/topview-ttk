package common

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

const aesKey = "3a5db96a8e844a6f444c1beac4b63a64e5c8773cfb862ea0516fbc3b197da21d"

var aesCrypto *AESCrypto

type AESCrypto struct {
	key []byte
}

func GetAESCrypto() *AESCrypto {
	once.Do(func() {
		aesCrypto = &AESCrypto{key: []byte(aesKey)}
	})
	return aesCrypto
}

func (ac *AESCrypto) Encrypt(plaintext []byte) ([]byte, error) {
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return append(nonce, ciphertext...), nil
}

func (ac *AESCrypto) Decrypt(ciphertext []byte) ([]byte, error) {
	block, err := aes.NewCipher(ac.key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plaintext, nil
}
