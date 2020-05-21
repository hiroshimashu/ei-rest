package main

import (
	"fmt"
	"log"
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

func NewInMemoryUserStore() *InMemoryUserStore {
	return &InMemoryUserStore{map[string]int{}}
}

type InMemoryUserStore struct {
	store map[string]int
}

func (i *InMemoryUserStore) GetUserId(id string) int {
	return i.store[id]
}

func (i *InMemoryUserStore) RecordUser(id string) {
	i.store[id]++
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

func main() {
	server := &UserServer{NewInMemoryUserStore()}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
