package user

type User struct {
	id   int
	guid string
}

func New(id int, guid string) User {
	return User{
		id:   id,
		guid: guid,
	}
}

func (user User) GetID() int {
	return user.id
}

func (user User) GetGUID() string {
	return user.guid
}

func (user *User) SetGUID(value string) {
	user.guid = value
}
