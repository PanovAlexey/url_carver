package user

type user struct {
	id   int
	guid string
}

func New(id int, guid string) user {
	return user{
		id:   id,
		guid: guid,
	}
}

func (user user) GetID() int {
	return user.id
}

func (user user) GetGUID() string {
	return user.guid
}

func (user *user) SetGuid(value string) {
	user.guid = value
}
