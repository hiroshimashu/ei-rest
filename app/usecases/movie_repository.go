package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MovieRepository interface {
	FindAll() (domain.Movies, error)
	FindByID(id string) (domain.Movie, error)
	Save(u domain.Movie) (domain.Movie, error)
}
