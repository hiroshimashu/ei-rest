package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	ids   map[string]int
	users []string
}

func (s *StubUserStore) GetUserId(user string) int {
	score := s.ids[user]
	return score
}

func (s *StubUserStore) RecordUser(id string) {
	s.users = append(s.users, id)
}

type StubMovie struct {
}

func TestGetUser(t *testing.T) {
	store := StubUserStore{
		map[string]int{
			"1": 1,
			"2": 2,
		},
		nil,
	}
	server := NewUserServer(&store)
	t.Run("returns user id 1", func(t *testing.T) {
		request := newGetUserIdRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "1")
	})

	t.Run("returns user id 2", func(t *testing.T) {
		request := newGetUserIdRequest("2")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "2")
	})

	t.Run("returns 404 on missing user", func(t *testing.T) {
		request := newGetUserIdRequest("3")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreUser(t *testing.T) {
	store := StubUserStore{
		map[string]int{},
		nil,
	}

	server := NewUserServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		user := "1"
		request := newPostUserRequest("1")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.users) != 1 {
			t.Errorf("got %d calls to RecordUsers want %d", len(store.users), 1)
		}

		if store.users[0] != user {
			t.Errorf("did't record correct user got %q want %q", store.users[0], user)
		}
	})
}

func TestRecordUserAndRetrievingThem(t *testing.T) {
	store := NewInMemoryUserStore()
	server := NewUserServer(store)
	user := "1"

	server.ServeHTTP(httptest.NewRecorder(), newPostUserRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostUserRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostUserRequest(user))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetUserIdRequest(user))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetUserIdRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", id), nil)
	return req
}

func newPostUserRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%s", id), nil)
	return req
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q want %q", got, want)
	}
}
