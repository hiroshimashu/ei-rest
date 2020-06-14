package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MyMovieRepository interface {
	FindMyMovies(userid string) (domain.MyMovies, error)
	Save(u domain.MyMovie) (domain.MyMovie, error)
}
