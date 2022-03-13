package services

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/hex"
)

type encryptor struct {
	gcm   cipher.AEAD
	nonce []byte
}

func NewEncryptionService() (*encryptor, error) {
	userKey := "23232323" //@ToDo
	key := sha256.Sum256([]byte(userKey))

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
