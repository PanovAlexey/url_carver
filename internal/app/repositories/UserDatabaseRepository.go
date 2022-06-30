package repositories

import (
	"database/sql"
	"github.com/PanovAlexey/url_carver/internal/app/domain/entity/user"
	"github.com/PanovAlexey/url_carver/internal/app/services/database"
)

type DatabaseUserRepository struct {
	DB *sql.DB
}

func GetDatabaseUserRepository(DB *sql.DB) *DatabaseUserRepository {
	return &DatabaseUserRepository{DB: DB}
}

func (repository DatabaseUserRepository) SaveUser(user user.User) (int, error) {
	var insertedID int

	query := "INSERT INTO " + database.TableUsersName + " (guid) VALUES ($1) RETURNING id"
	err := repository.DB.QueryRow(query, user.GetGUID()).Scan(&insertedID)

	if err != nil {
		return 0, err // ToDo: 0 - is a crutch
	}

	return insertedID, err
}

func (repository DatabaseUserRepository) GetUserByGUID(guid string) (user.User, error) {
	query := "SELECT id FROM " + database.TableUsersName + " WHERE guid=($1)"
	row := repository.DB.QueryRow(query, guid)

	var userID int
	err := row.Scan(&userID)

	if err != nil {
		return user.User{}, err
	}

	return user.New(userID, guid), nil
}

func (repository DatabaseUserRepository) GetUserByID(userID int) (user.User, error) {
	query := "SELECT id FROM " + database.TableUsersName + " WHERE id=($1)"
	row := repository.DB.QueryRow(query, userID)

	var userGUID string
	err := row.Scan(&userGUID)

	if err != nil {
		return user.User{}, err
	}

	return user.New(userID, userGUID), nil
}

func (repository DatabaseUserRepository) IsExistUserByGUID(guid string) bool {
	user, err := repository.GetUserByGUID(guid)

	if err != nil || user.GetID() < 1 {
		return false
	}

	return true
}
