package usecases

import (
	"github.com/hiroshimashu/ei-rest/app/domain"
)

type MockMyMovieInteractor struct {
	MyMovieRepository MyMovieRepository
}

func (mi *MockMyMovieInteractor) Index(userid string) (mymovies domain.MyMovies, err error) {
	mymovies, err = mi.MyMovieRepository.FindMyMovies(userid)
	return
}

func (mi *MockMyMovieInteractor) Create(newMyMovie domain.MyMovie) (mymovie domain.MyMovie, err error) {
	mymovie, err = mi.MyMovieRepository.Save(newMyMovie)
	return
}

func NewMockMyMovieInteractor(mr *MockMyMovieRepository) *MockMyMovieInteractor {
	return &MockMyMovieInteractor{
		MyMovieRepository: mr,
	}
}

type MockMyMovieRepository struct {
}

func (mr *MockMyMovieRepository) FindMyMovies(userid string) (mymovies domain.MyMovies, err error) {
	mymovies, err = mr.FindMyMovies(userid)
	return
}

func (mr *MockMyMovieRepository) Save(newMyMovie domain.MyMovie) (mymovie domain.MyMovie, err error) {
	mymovie, err = mr.Save(newMyMovie)
	return
}
