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
	GetUserByID(userID int) (user.UserInterface, error)
	GetUserByGUID(guid string) (user.UserInterface, error)
	IsExistUserByGUID(guid string) bool
}

func GetDatabaseUserService(
	databaseRepository databaseUserRepositoryInterface,
) *databaseUserService {
	return &databaseUserService{databaseRepository: databaseRepository}
}

func (service databaseUserService) SaveUser(user user.UserInterface) (int, error) {
	insertedID, err := service.databaseRepository.SaveUser(user)

	if err != nil {
		log.Println("user dit not save to database: " + err.Error())
	}

	return insertedID, err
}

func (service databaseUserService) GetUserByToken(token string) (user.UserInterface, error) {
	return service.databaseRepository.GetUserByGUID(token)
}

func (service databaseUserService) GetUserByID(userID int) (user.UserInterface, error) {
	return service.databaseRepository.GetUserByID(userID)
}

func (service databaseUserService) IsExistUserByToken(token string) bool {
	return service.databaseRepository.IsExistUserByGUID(token)
}
