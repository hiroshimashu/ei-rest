package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MovieInteractor struct {
	MovieRepository MovieRepository
}

func (MovieInteractor *MovieInteractor) Index() (movies domain.Movies, err error) {
	movies, err = MovieInteractor.MovieRepository.FindAll()
	return
}

func (MovieInteractor *MovieInteractor) IndexByID(id string) (movie domain.Movie, err error) {
	movie, err = MovieInteractor.MovieRepository.FindByID(id)
	return
}

func (MovieInteractor *MovieInteractor) Store(newMovie domain.Movie) (movie domain.Movie, err error) {
	movie, err = MovieInteractor.MovieRepository.Save(newMovie)
	return
}
