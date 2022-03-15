package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
	"github.com/PanovAlexey/url_carver/config"
)

type encryptor struct {
	gcm   cipher.AEAD
	nonce []byte
}

func NewEncryptionService(config config.Config) (*encryptor, error) {
	key := sha256.Sum256([]byte(config.GetEncryptionKey()))

	block, err := aes.NewCipher(key[:])

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	nonce := key[len(key)-gcm.NonceSize():]

	return &encryptor{
		gcm:   gcm,
		nonce: nonce,
	}, nil
}

func (encryptor *encryptor) Encrypt(data string) string {
	encryptedBytes := encryptor.gcm.Seal(nil, encryptor.nonce, []byte(data), nil)

	return hex.EncodeToString(encryptedBytes)
}

func (encryptor *encryptor) Decrypt(msg string) (string, error) {
	msgBytes, err := hex.DecodeString(msg)

	if err != nil {
		return "", err
	}

	decryptedBytes, err := encryptor.gcm.Open(nil, encryptor.nonce, msgBytes, nil)

	if err != nil {
		return "", err
	}

	return string(decryptedBytes), nil
}
