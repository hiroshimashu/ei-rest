package infrastructure

import (
	"app/domain/repository"
	"app/domain/model"
)

type UserRepository struct {
	SqlHandler
}

func NewUserRepository(sqlHandler SqlHandler) repository.UserRepository {
	userRepository := UserRepository{sqlHandler}
	return &UserRepository
}

func (ur *UserRepository) FindAll() (users []*model.User, error) {

}

func (ur *UserRepository) Find(id string) (user *model.User, error) {

}

func (ur *UserRepository) Create(user *model.User) (user *model.User, error) {

}

func (ur *UserRepository) Update(user *model.User) (user *model.User, error) {

}
