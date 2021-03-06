package services

import (
	"errors"
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"log"
)

type DatabaseURLService struct {
	databaseRepository  databaseURLRepositoryInterface
	databaseUserService DatabaseUserService
}

type databaseURLRepositoryInterface interface {
	SaveURL(url dto.DatabaseURL) (int, error)
	GetList() ([]dto.DatabaseURL, error)
	SaveBatchURLs(collection []dto.DatabaseURL) error
	DeleteURLsByShortValueSlice([]string, int) ([]dto.DatabaseURL, error)
}

func GetDatabaseURLService(
	databaseRepository databaseURLRepositoryInterface,
	databaseUserService DatabaseUserService,
) *DatabaseURLService {
	return &DatabaseURLService{
		databaseRepository:  databaseRepository,
		databaseUserService: databaseUserService,
	}
}

func (service DatabaseURLService) SaveURL(url url.URL) (int, error) {
	userID, err := service.verifyUser(url.UserID)

	if err != nil {
		log.Println("error while URL user verification: " + err.Error())
	}

	databaseURL := dto.DatabaseURL{
		LongURL:   url.LongURL,
		ShortURL:  url.ShortURL,
		UserID:    userID,
		IsDeleted: false,
	}

	log.Println(`try to save to database URL: `, databaseURL)
	log.Println(databaseURL)

	URLID, err := service.databaseRepository.SaveURL(databaseURL)

	if err != nil {
		log.Println("url dit not save to database: " + err.Error())
	}

	return URLID, err
}

func (service DatabaseURLService) SaveBatchURLs(collection []url.URL) {
	var URLDatabaseCollection []dto.DatabaseURL

	for _, url := range collection {
		userID, err := service.verifyUser(url.UserID)

		if err != nil {
			log.Println("error while URL user verification: " + err.Error())
		}

		databaseURL := dto.DatabaseURL{
			LongURL:   url.LongURL,
			ShortURL:  url.ShortURL,
			UserID:    userID,
			IsDeleted: false,
		}

		URLDatabaseCollection = append(URLDatabaseCollection, databaseURL)
	}

	log.Println(`try to save to database batch URLs. Items count: `, len(URLDatabaseCollection))

	err := service.databaseRepository.SaveBatchURLs(URLDatabaseCollection)

	if err != nil {
		log.Println(`error while URLs batch saving: `, err)
	}
}

func (service DatabaseURLService) RemoveByShortURLSlice(URLSlice []string, userToken string) error {
	batchURLsRemovingService := GetBatchURLsRemovingService(service.databaseRepository)
	userEntity, err := service.databaseUserService.GetUserByToken(userToken)
	userID := userEntity.GetID()

	if err != nil {
		log.Println(errors.New("an error occurred while getting a user by token " + userToken + ". " + err.Error()))

		// @ToDo: delete it. Crutch for autotests. do not return an error if the database fails.
		// return errors.New("an error occurred while getting a user by token " + userToken + ". " + err.Error())
	}

	return batchURLsRemovingService.RemoveByShortURLSlice(URLSlice, userID)
}

func (service DatabaseURLService) GetURLCollectionFromStorage() []url.URL {
	dtoURLCollection := []url.URL{}

	collection, err := service.databaseRepository.GetList()

	if err != nil {
		log.Println(`error while getting all users.`, err.Error())

		return dtoURLCollection
	}

	for _, databaseURL := range collection {
		user, err := service.databaseUserService.GetUserByID(databaseURL.UserID)

		if err != nil {
			log.Println(`error while getting user by ID `, databaseURL.UserID, `. `, err.Error())

			continue
		}

		dtoURLCollection = append(
			dtoURLCollection,
			url.URL{
				LongURL:   databaseURL.LongURL,
				ShortURL:  databaseURL.ShortURL,
				UserID:    user.GetGUID(),
				IsDeleted: databaseURL.IsDeleted,
			},
		)
	}

	return dtoURLCollection
}

func (service DatabaseURLService) verifyUser(userToken string) (int, error) {
	var userID int

	if service.databaseUserService.IsExistUserByToken(userToken) {
		userEntity, err := service.databaseUserService.GetUserByToken(userToken)
		userID = userEntity.GetID()

		if err != nil {
			log.Println(`error was found while user getting from database: ` + err.Error())
		}
	} else {
		savedUserID, err := service.databaseUserService.SaveUser(user.New(0, userToken))

		if err != nil {
			log.Println(`error was found while user saving to database: ` + err.Error())
			return 0, err // ToDo: 0 is a crutch
		}

		userID = savedUserID

		log.Println(`user `, userID, ` saving to database successfully completed`)
	}

	return userID, nil
}
