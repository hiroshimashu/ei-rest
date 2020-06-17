package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MyMovieRepository interface {
	FindMyMovies(userid string) (domain.MyMovies, error)
	Save(mymovie domain.MyMovie) (domain.MyMovie, error)
}
