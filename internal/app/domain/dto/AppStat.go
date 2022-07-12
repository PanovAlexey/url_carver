package dto

type AppStat struct {
	Urls  int `json:"urls"`
	Users int `json:"users"`
}

func GetAppStatByURLsCountAndUsersCount(URLsCount int, usersCount int) AppStat {
	return AppStat{Urls: URLsCount, Users: usersCount}
}
