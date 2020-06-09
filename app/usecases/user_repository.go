package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type UserRepository interface {
	FindAll() (domain.Users, error)
	Store(u domain.User) (domain.User, error)
}
