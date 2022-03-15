package encryption

type EncryptorInterface interface {
	Encrypt(data string) string
	Decrypt(encryptedData string) (string, error)
}
