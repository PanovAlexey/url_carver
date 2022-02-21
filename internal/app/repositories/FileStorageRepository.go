package repositories

import (
	"bufio"
	"os"
)

const filePermissions = 0777

type fileStorageRepository struct {
	config  FileStorageConfigInterface
	writer  *bufio.Writer
	scanner *bufio.Scanner
}

type FileStorageConfigInterface interface {
	GetFileStoragePath() string
}

func GetFileStorageRepository(config FileStorageConfigInterface) (*fileStorageRepository, error) {
	file, err := os.OpenFile(config.GetFileStoragePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, filePermissions)

	if err != nil {
		return nil, err
	}

	return &fileStorageRepository{
		writer:  bufio.NewWriter(file),
		scanner: bufio.NewScanner(file),
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

func (repository *fileStorageRepository) ReadLine() ([]byte, error) {
	if !repository.scanner.Scan() {
		return nil, repository.scanner.Err()
	}

	return repository.scanner.Bytes(), nil
}
