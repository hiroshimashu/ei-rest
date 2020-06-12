package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hiroshimashu/ei-rest/app/domain"
	"github.com/hiroshimashu/ei-rest/app/usecases"
)

type MockMovieConroller struct {
	Interactor *usecases.MockMovieInteractor
}

func NewMovieController() *MockMovieConroller {
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
	mr := usecases.NewMockRepository(movies)
	mi := usecases.NewMockMovieInteractor(mr)
	return &MockMovieConroller{
		Interactor: mi,
	}
}

func (mc *MockMovieConroller) Index(w http.ResponseWriter, r *http.Request) {
	movies, err := mc.Interactor.Index()
	if err != nil {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func (mc *MockMovieConroller) Create(w http.ResponseWriter, r *http.Request) {
	var movie domain.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	storedMovie, err := mc.Interactor.Store(movie)
	if err != nil {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(storedMovie)
}
