package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type UserInteractor struct {
	UserRepository UserRepository
}

func (ui *UserInteractor) Index() (users domain.Users, err error) {
	users, err = ui.UserRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ui *UserInteractor) Add(u domain.User) (domain.User, error) {
	user, err := ui.UserRepository.Store(u)
	return user, err
}
