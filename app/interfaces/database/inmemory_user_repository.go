package database

import "github.com/hiroshimashu/ei-rest/app/domain"

func NewImMemoryUserRepositry() *InMemoryUserRepository {
	return &InMemoryUserRepository{}
}

type InMemoryUserRepository struct {
}

func (ur *InMemoryUserRepository) FindAll() (domain.Users, error) {
	users := []domain.User{
		{
			ID:       "5555",
			Gender:   "M",
			Age:      26,
			Area:     "tokyo",
			Email:    "example1@g.com",
			Password: "1111",
			IsPaid:   true,
		},
		{
			ID:       "4444",
			Gender:   "W",
			Age:      28,
			Area:     "kyoto",
			Email:    "example2@g.com",
			Password: "2222",
			IsPaid:   false,
		},
	}

	return users, nil
}

func (ur *InMemoryUserRepository) Store(u domain.User) (domain.User, error) {

	return u, nil
}
