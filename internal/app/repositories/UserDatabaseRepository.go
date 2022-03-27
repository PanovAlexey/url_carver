package repositories

import (
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
)

type databaseUserRepository struct {
	databaseService database.DatabaseInterface
}

func GetDatabaseUserRepository(databaseService database.DatabaseInterface) *databaseUserRepository {
	return &databaseUserRepository{databaseService: databaseService}
}

func (repository databaseUserRepository) SaveUser(user user.UserInterface) (int, error) {
	var insertedID int

	query := "INSERT INTO " + database.TableUsersName + " (guid) VALUES ($1) RETURNING id"
	err := repository.databaseService.GetDatabaseConnection().
		QueryRow(query, user.GetGUID()).Scan(&insertedID)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedID, err
}

func (repository databaseUserRepository) GetUserByGUID(guid string) (user.UserInterface, error) {
	query := "SELECT id FROM " + database.TableUsersName + " WHERE guid=($1)"
	row := repository.databaseService.GetDatabaseConnection().QueryRow(query, guid)

	var userID int
	err := row.Scan(&userID)

	if err != nil {
		return nil, err
	}

	return user.New(userID, guid), nil
}

func (repository databaseUserRepository) GetUserByID(userID int) (user.UserInterface, error) {
	query := "SELECT id FROM " + database.TableUsersName + " WHERE id=($1)"
	row := repository.databaseService.GetDatabaseConnection().QueryRow(query, userID)

	var userGUID string
	err := row.Scan(&userGUID)

	if err != nil {
		return nil, err
	}

	return user.New(userID, userGUID), nil
}

func (repository databaseUserRepository) IsExistUserByGUID(guid string) bool {
	user, err := repository.GetUserByGUID(guid)

	if err != nil || user.GetID() < 1 {
		return false
	}

	return true
}
