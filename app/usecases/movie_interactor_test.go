package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/domain"
)

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
	movie, err = mockMovieInteractor.MovieRepository.Save(newMovie)
	return
}

func NewMockMovieInteractor(mc *MockMovieRepository) *MockMovieInteractor {
	return &MockMovieInteractor{
		MovieRepository: mc,
	}
}

type MockMovieRepository struct {
	Movies domain.Movies
}

func (mr *MockMovieRepository) FindAll() (domain.Movies, error) {
	return mr.Movies, nil
}

func (mr *MockMovieRepository) FindByID(id string) (domain.Movie, error) {
	for _, v := range mr.Movies {
		if v.ID == id {
			return v, nil
		}
	}
	return domain.Movie{}, nil
}

func (mr *MockMovieRepository) Save(movie domain.Movie) (domain.Movie, error) {
	mr.Movies = append(mr.Movies, movie)
	return movie, nil
}

func TestMovieInteractor(t *testing.T) {
	mr := &MockMovieRepository{
		Movies: domain.Movies{
			domain.Movie{
				ID:  "3333",
				URL: "https://example.com",
			},
			domain.Movie{
				ID:  "5555",
				URL: "https://example2.com",
			},
		},
	}
	mockMovieInteractor := NewMockMovieInteractor(mr)

	t.Run("Correctly Index method works", func(t *testing.T) {
		got, err := mockMovieInteractor.Index()
		if err != nil {
			t.Errorf("Get movie has failed")
		}
		want := domain.Movies{
			domain.Movie{
				ID:  "3333",
				URL: "https://example.com",
			},
			domain.Movie{
				ID:  "5555",
				URL: "https://example2.com",
			},
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Correctly Store method works", func(t *testing.T) {
		newMovie := domain.Movie{
			ID:  "77777",
			URL: "https://example3.com",
		}

		got, err := mockMovieInteractor.Store(newMovie)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(mockMovieInteractor.Index())
		want := newMovie
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}

	})
}
