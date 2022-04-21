package services

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/dto"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/url"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"log"
)

type databaseURLService struct {
	databaseRepository  databaseURLRepositoryInterface
	databaseUserService databaseUserService
}

type databaseURLRepositoryInterface interface {
	SaveURL(url dto.DatabaseURL) (int, error)
	GetList() ([]dto.DatabaseURL, error)
	SaveBatchURLs(collection []dto.DatabaseURL) error
	DeleteURLsByShortValueSlice([]string, int) ([]dto.DatabaseURL, error)
}

func GetDatabaseURLService(
	databaseRepository databaseURLRepositoryInterface,
	databaseUserService databaseUserService,
) *databaseURLService {
	return &databaseURLService{
		databaseRepository:  databaseRepository,
		databaseUserService: databaseUserService,
	}
}

func (service databaseURLService) SaveURL(url url.URL) (int, error) {
	userID, err := service.verifyUser(url.GetUserToken())

	if err != nil {
		log.Println("error while URL user verification: " + err.Error())
	}

	databaseURL := dto.NewDatabaseURL(
		url.GetLongURL(),
		url.GetShortURL(),
		userID,
		false,
	)

	log.Println(`try to save to database URL: `, databaseURL)
	log.Println(databaseURL)

	URLID, err := service.databaseRepository.SaveURL(databaseURL)

	if err != nil {
		log.Println("url dit not save to database: " + err.Error())
	}

	return URLID, err
}

func (service databaseURLService) SaveBatchURLs(collection []url.URL) {
	var URLDatabaseCollection []dto.DatabaseURL

	for _, url := range collection {
		userID, err := service.verifyUser(url.GetUserToken())

		if err != nil {
			log.Println("error while URL user verification: " + err.Error())
		}

		URLDatabaseCollection = append(
			URLDatabaseCollection,
			dto.NewDatabaseURL(url.GetLongURL(), url.GetShortURL(), userID, false),
		)
	}

	log.Println(`try to save to database batch URLs. Items count: `, len(URLDatabaseCollection))

	err := service.databaseRepository.SaveBatchURLs(URLDatabaseCollection)

	if err != nil {
		log.Println(`error while URLs batch saving: `, err)
	}
}

func (service databaseURLService) RemoveByShortURLSlice(URLSlice []string, userToken string) error {
	batchURLsRemovingService := GetBatchURLsRemovingService(service.databaseRepository)
	userEntity, err := service.databaseUserService.GetUserByToken(userToken)
	userID := userEntity.GetID()

	if err != nil {
		return err
	}

	return batchURLsRemovingService.RemoveByShortURLSlice(URLSlice, userID)
}

func (service databaseURLService) verifyUser(userToken string) (int, error) {
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

func (service databaseURLService) GetURLCollectionFromStorage() []url.URL {
	dtoURLCollection := []url.URL{}

	collection, err := service.databaseRepository.GetList()

	if err != nil {
		log.Println(`error while getting all users.`, err.Error())

		return dtoURLCollection
	}

	for _, databaseURL := range collection {
		user, err := service.databaseUserService.GetUserByID(databaseURL.GetUserID())

		if err != nil {
			log.Println(`error while getting user by ID `, databaseURL.GetUserID(), `. `, err.Error())

			continue
		}

		dtoURLCollection = append(
			dtoURLCollection,
			url.New(databaseURL.GetLongURL(), databaseURL.GetShortURL(), user.GetGUID(), databaseURL.GetIsDeleted()),
		)
	}

	return dtoURLCollection
}
