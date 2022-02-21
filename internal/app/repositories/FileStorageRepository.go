package repositories

import (
	"bufio"
	"os"
)

const filePermissions = 0777

type fileStorageRepository struct {
	config FileStorageConfigInterface
	writer *bufio.Writer
}

type FileStorageConfigInterface interface {
	GetFileStoragePath() string
}

func GetFileStorageRepository(config FileStorageConfigInterface) (*fileStorageRepository, error) {
	file, err := os.OpenFile(config.GetFileStoragePath(), os.O_WRONLY|os.O_CREATE|os.O_APPEND, filePermissions)

	if err != nil {
		return nil, err
	}

	return &fileStorageRepository{
		writer: bufio.NewWriter(file),
	}, nil
}

func (repository *fileStorageRepository) WriteLine(data []byte) error {
	if _, err := repository.writer.Write(data); err != nil {
		return err
	}

	if err := repository.writer.WriteByte('\n'); err != nil {
		return err
	}

	return repository.writer.Flush()
}
