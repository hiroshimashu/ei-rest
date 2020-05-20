package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	ids map[string]int
}

func (s *StubUserStore) GetUserId(user string) int {
	score := s.ids[user]
	return score
}

func TestGetUser(t *testing.T) {
	store := StubUserStore{
		map[string]int{
			"1": 1,
			"2": 2,
		},
	}
	server := &UserServer{&store}
	t.Run("returns user id 1", func(t *testing.T) {
		request := newGetUserIdRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "1")
	})

	t.Run("returns user id 2", func(t *testing.T) {
		request := newGetUserIdRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertResponseBody(t, response.Body.String(), "2")
	})
}

func newGetUserIdRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
