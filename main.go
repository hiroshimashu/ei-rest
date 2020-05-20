package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserId(id string) int
}

type UserServer struct {
	store UserStore
}

func (u *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Fprint(w, u.store.GetUserId(user))
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
	server := &UserServer{}
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
