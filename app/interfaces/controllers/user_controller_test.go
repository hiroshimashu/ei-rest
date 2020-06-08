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

func TestUserRouter(t *testing.T) {
	t.Run("Correctly returns users", func(t *testing.T) {
		mockUserController := NewMockUserController()
		reqBody := bytes.NewBufferString("")
		req := httptest.NewRequest(http.MethodGet, "/users", reqBody)

		res := httptest.NewRecorder()
		mockUserController.Index(res, req)

		assertStatus(t, res.Code, http.StatusOK)

		got := getUsersFromResponse(t, res.Body)

		want := []domain.User{
			{
				ID:       "5555",
				Gender:   "M",
				Age:      26,
				Area:     "tokyo",
				Email:    "example1@g.com",
				Password: "1111",
				IsPaid:   true,
			},
			{
				ID:       "4444",
				Gender:   "W",
				Age:      28,
				Area:     "kyoto",
				Email:    "example2@g.com",
				Password: "2222",
				IsPaid:   false,
			},
		}

		assertUsers(t, got, want)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status got %d, want %d", got, want)
	}
}

func getUsersFromResponse(t *testing.T, body io.Reader) (users []domain.User) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&users)

	if err != nil {
		t.Fatal("unable to parse response from server")
	}

	return
}

func assertUsers(t *testing.T, got, want []domain.User) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
