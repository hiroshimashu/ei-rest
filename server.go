package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	id  string
	url string
}

type UserStore interface {
	GetUserId(id string) int
	RecordUser(id string)
}

type UserServer struct {
	store UserStore
	http.Handler
}

func NewUserServer(store UserStore) *UserServer {
	u := new(UserServer)

	u.store = store

	router := http.NewServeMux()
	router.Handle("/users/", http.HandlerFunc(u.usersHandler))
	router.Handle("/movies/", http.HandlerFunc(u.moviesHanlder))

	u.Handler = router

	return u
}

func (u *UserServer) moviesHanlder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(u.getMovieTable())
}

func (u *UserServer) getMovieTable() []Movie {
	return []Movie{
		{"1111", "http://movie.com/1111"},
	}
}

func (u *UserServer) usersHandler(w http.ResponseWriter, r *http.Request) {
	user := r.URL.Path[len("/users/"):]

	switch r.Method {
	case http.MethodPost:
		u.processUser(w, user)
	case http.MethodGet:
		u.showId(w, user)
	}
}

func (u *UserServer) showId(w http.ResponseWriter, user string) {
	userId := u.store.GetUserId(user)

	if userId == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, userId)

}

func (u *UserServer) processUser(w http.ResponseWriter, user string) {
	u.store.RecordUser(user)
	w.WriteHeader(http.StatusAccepted)
}

func GetUserId(id string) string {
	if id == "1" {
		return "1"
	}

	if id == "2" {
		return "2"
	}

	return ""
}
