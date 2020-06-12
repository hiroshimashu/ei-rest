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
	t.Run("Correctly get movies", func(t *testing.T) {
		mockMovieController := NewMockMovieController()
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
}

func getMoviesFromResponse(t *testing.T, body io.Reader) (movies domain.Movies) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&movies)

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
