package usecases

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/domain"
)

func TestMovieInteractor(t *testing.T) {
	movies := domain.Movies{
		domain.Movie{
			ID:  "3333",
			URL: "https://example.com",
		},
		domain.Movie{
			ID:  "5555",
			URL: "https://example2.com",
		},
	}
	mr := NewMockRepository(movies)
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

		assertMovies(t, got, want)
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

		assertMovie(t, got, want)

	})

	t.Run("Correctly IndexByID method works", func(t *testing.T) {
		id := "5555"
		got, err := mockMovieInteractor.IndexByID(id)
		if err != nil {
			t.Error(err)
		}
		want := domain.Movie{
			ID:  "5555",
			URL: "https://example2.com",
		}

		assertMovie(t, got, want)

	})
}

func assertMovie(t *testing.T, got, want domain.Movie) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertMovies(t *testing.T, got, want domain.Movies) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
