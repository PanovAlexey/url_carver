package repositories

import (
	"bufio"
	"github.com/PanovAlexey/url_carver/config"
	"os"
)

const filePermissions = 0777

type fileStorageRepository struct {
	file    *os.File
	writer  *bufio.Writer
	scanner *bufio.Scanner
}

func GetFileStorageRepository(config config.Config) (*fileStorageRepository, error) {
	file, err := os.OpenFile(config.GetFileStoragePath(), os.O_RDWR|os.O_CREATE|os.O_APPEND, filePermissions)

	return &fileStorageRepository{
		writer:  bufio.NewWriter(file),
		scanner: bufio.NewScanner(file),
		file:    file,
	}, err
}

func (repository *fileStorageRepository) IsStorageExist() (bool, error) {
	isStorageExist := repository.file != nil

	return isStorageExist, nil
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

func (repository *fileStorageRepository) Close() {
	repository.file.Close()
}
