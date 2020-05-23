package main

import (
	"fmt"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserId(id string) int
	RecordUser(id string)
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")
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
