package usecases

import "github.com/hiroshimashu/ei-rest/app/domain"

type MockMovieInteractor struct {
	MovieRepository MovieRepository
}

func (mockMovieInteractor *MockMovieInteractor) Index() (movies domain.Movies, err error) {
	movies, err = mockMovieInteractor.MovieRepository.FindAll()
	return
}

func (mockMovieInteractor *MockMovieInteractor) IndexByID(id string) (movie domain.Movie, err error) {
	movie, err = mockMovieInteractor.MovieRepository.FindByID(id)
	return
}

func (mockMovieInteractor *MockMovieInteractor) Store(newMovie domain.Movie) (movie domain.Movie, err error) {
	movie, err = mockMovieInteractor.MovieRepository.Save(movie)
	return
}
