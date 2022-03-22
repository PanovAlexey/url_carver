package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"log"
)

type databaseUserService struct {
	databaseRepository databaseUserRepositoryInterface
}

type databaseUserRepositoryInterface interface {
	SaveUser(user user.UserInterface) (int, error)
	GetUserByGuid(guid string) (user.UserInterface, error)
	IsExistUserByGuid(guid string) bool
}

func GetDatabaseUserService(
	databaseRepository databaseUserRepositoryInterface,
) *databaseUserService {
	return &databaseUserService{databaseRepository: databaseRepository}
}

func (service databaseUserService) SaveUser(user user.UserInterface) (int, error) {
	insertedId, err := service.databaseRepository.SaveUser(user)

	if err != nil {
		log.Println("user dit not save to database: " + err.Error())
	}

	return insertedId, err
}

func (service databaseUserService) GetUserByToken(token string) (user.UserInterface, error) {
	return service.databaseRepository.GetUserByGuid(token)
}

func (service databaseUserService) IsExistUserByToken(token string) bool {
	return service.databaseRepository.IsExistUserByGuid(token)
}
