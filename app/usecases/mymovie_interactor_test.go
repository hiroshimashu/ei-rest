package usecases

import (
	"reflect"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/domain"
)

func TestMyMovieInteractor(t *testing.T) {
	mymovies := domain.MyMovies{
		domain.MyMovie{
			ID:      "111111",
			UserID:  "5555",
			MovieID: "6666",
		},
		domain.MyMovie{
			ID:      "222222",
			UserID:  "5555",
			MovieID: "8888",
		},
	}
	mr := MockMyMovieRepository{
		MyMovies: mymovies,
	}
	mockMyMovieInteractor := NewMockMyMovieInteractor(&mr)

	t.Run("Correctly return mymovies", func(t *testing.T) {
		got, err := mockMyMovieInteractor.Index("5555")
		if err != nil {
			t.Error(err)
		}
		want := domain.MyMovies{
			domain.MyMovie{
				ID:      "111111",
				UserID:  "5555",
				MovieID: "6666",
			},
			domain.MyMovie{
				ID:      "222222",
				UserID:  "5555",
				MovieID: "8888",
			},
		}

		assertMyMovies(t, got, want)
	})

	t.Run("Correctly return empty mymovies when userid is not matched", func(t *testing.T) {
		got, err := mockMyMovieInteractor.Index("7777")
		if err != nil {
			t.Error(err)
		}
		want := domain.MyMovies{}

		assertMyMovies(t, got, want)
	})

	t.Run("Correctly inserting mymovie works", func(t *testing.T) {
		newMovie := domain.MyMovie{
			ID:      "222222",
			UserID:  "7777",
			MovieID: "9999",
		}
		err := mockMyMovieInteractor.Create(newMovie)

		if err != nil {
			t.Error(err)
		}

	})
}

func assertMyMovie(t *testing.T, got, want domain.MyMovie) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertMyMovies(t *testing.T, got, want domain.MyMovies) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
