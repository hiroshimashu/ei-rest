package repository

import (
	"app/domain/model"
)

type UserRepository interface {
	FindAll(users []*model.User)
	Find(id string) (user *model.User, err error)
	Create(user *model.User) (*model.User, error)
	Update(user *model.User) (*model.User, error)
}
