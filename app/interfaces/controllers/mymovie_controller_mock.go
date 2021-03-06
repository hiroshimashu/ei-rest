package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/hiroshimashu/ei-rest/app/domain"
	"github.com/hiroshimashu/ei-rest/app/usecases"
)

type MockMyMovieController struct {
	Interactor *usecases.MockMyMovieInteractor
}

func (mc *MockMyMovieController) Index(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Path
	sk := strings.Split(keys, "/")
	id := sk[2]
	if id == " " {
		log.Fatal("No id has given")
	}
	mymovies, err := mc.Interactor.Index(id)
	if err != nil {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mymovies)
}

func (mc *MockMyMovieController) Store(w http.ResponseWriter, r *http.Request) {
	var mymovie domain.MyMovie
	err := json.NewDecoder(r.Body).Decode(&mymovie)

	if err != nil {
		log.Fatal(err)
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	err = mc.Interactor.Create(mymovie)

	if err != nil {
		log.Fatal(err)
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
}

func NewMockMyMovieController() *MockMyMovieController {
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
	mr := usecases.MockMyMovieRepository{
		MyMovies: mymovies,
	}
	mi := usecases.NewMockMyMovieInteractor(&mr)
	return &MockMyMovieController{
		Interactor: mi,
	}
}
