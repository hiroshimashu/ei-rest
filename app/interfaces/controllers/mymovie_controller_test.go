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

func TestMyMovieController(t *testing.T) {
	mockMyMovieController := NewMockMyMovieController()

	t.Run("Correctly get mymovies", func(t *testing.T) {
		reqBody := bytes.NewBufferString("5555")
		req := httptest.NewRequest(http.MethodGet, "/mymovies", reqBody)
		res := httptest.NewRecorder()

		mockMyMovieController.Index(res, req)
		AssertStatus(t, res.Code, http.StatusOK)

		got := getMyMoviesFromResponse(t, res.Body)
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
}

func getMyMoviesFromResponse(t *testing.T, body io.Reader) (movie domain.MyMovies) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&movie)

	if err != nil {
		t.Fatal("unable to parse response from server")
	}

	return
}

func assertMyMovies(t *testing.T, got, want domain.MyMovies) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
