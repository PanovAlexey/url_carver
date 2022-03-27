package repositories

type FileStorageRepositoryInterface interface {
	IsStorageExist() (bool, error)
	WriteLine(data []byte) error
	ReadLine() ([]byte, error)
	Close()
}
