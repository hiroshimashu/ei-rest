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
	MyMovies domain.MyMovies
}

func (mr *MockMyMovieRepository) FindMyMovies(userid string) (mymovies domain.MyMovies, err error) {
	movies := Filter(mr.MyMovies, func(mymovie domain.MyMovie) bool {
		return mymovie.UserID == userid
	})
	return movies, nil
}

func (mr *MockMyMovieRepository) Save(newMyMovie domain.MyMovie) (mymovie domain.MyMovie, err error) {
	mymovie, err = mr.Save(newMyMovie)
	return
}

func Filter(mymovies domain.MyMovies, f func(domain.MyMovie) bool) []domain.MyMovie {
	movies := make([]domain.MyMovie, 0)
	for _, v := range mymovies {
		if f(v) {
			movies = append(movies, v)
		}
	}
	return movies
}
