package usecases

import (
	"github.com/hiroshimashu/ei-rest/app/domain"
)

type UserRepository interface {
	FindAll(users []domain.User, err error)
}
