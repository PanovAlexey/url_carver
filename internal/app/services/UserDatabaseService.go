package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"github.com/PanovAlexey/url_carver/internal/app/repositories"
	"log"
)

type DatabaseUserService struct {
	databaseRepository repositories.DatabaseUserRepository
}

func GetDatabaseUserService(
	databaseRepository repositories.DatabaseUserRepository,
) *DatabaseUserService {
	return &DatabaseUserService{databaseRepository: databaseRepository}
}

func (service DatabaseUserService) SaveUser(user user.User) (int, error) {
	insertedID, err := service.databaseRepository.SaveUser(user)

	if err != nil {
		log.Println("user dit not save to database: " + err.Error())
	}

	return insertedID, err
}

func (service DatabaseUserService) GetUserByToken(token string) (user.User, error) {
	return service.databaseRepository.GetUserByGUID(token)
}

func (service DatabaseUserService) GetUserByID(userID int) (user.User, error) {
	return service.databaseRepository.GetUserByID(userID)
}

func (service DatabaseUserService) IsExistUserByToken(token string) bool {
	return service.databaseRepository.IsExistUserByGUID(token)
}
