package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/domain"
)

func TestMovieController(t *testing.T) {
	mockMovieController := NewMockMovieController()
	t.Run("Correctly get movies", func(t *testing.T) {
		reqBody := bytes.NewBufferString("")
		req := httptest.NewRequest(http.MethodGet, "/movies", reqBody)

		res := httptest.NewRecorder()

		mockMovieController.Index(res, req)
		AssertStatus(t, res.Code, http.StatusOK)

		got := getMoviesFromResponse(t, res.Body)

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

	t.Run("Correctly post movie", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"ID":"7777","URL":"https://example3.com"}`)
		req := httptest.NewRequest(http.MethodPost, "/movies", reqBody)

		res := httptest.NewRecorder()

		mockMovieController.Create(res, req)
		AssertStatus(t, res.Code, http.StatusOK)

		got := getMovieFromResponse(t, res.Body)
		want := domain.Movie{
			ID:  "7777",
			URL: "https://example3.com",
		}

		assertMovie(t, got, want)
	})
}

func getMoviesFromResponse(t *testing.T, body io.Reader) (movies domain.Movies) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&movies)

	if err != nil {
		t.Fatal("unable to parse response from server")
	}

	return
}

func getMovieFromResponse(t *testing.T, body io.Reader) (movie domain.Movie) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&movie)

	if err != nil {
		t.Fatal("unable to parse response from server")
	}

	return
}

func assertMovies(t *testing.T, got, want domain.Movies) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertMovie(t *testing.T, got, want domain.Movie) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
