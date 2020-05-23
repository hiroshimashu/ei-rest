package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMovies(t *testing.T) {
	store := StubUserStore{}
	server := NewUserServer(&store)

	t.Run("it return 200 on /movies", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/movies/", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		var got []Movie

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of User, '%v'", response.Body, err)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})
}
