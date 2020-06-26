package infrastructure

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hiroshimashu/ei-rest/app/interfaces/controllers"
)

func TestUserRouter(t *testing.T) {
	server := NewMockServer()

	t.Run("it returns 200 on /users", func(t *testing.T) {
		reqBody := bytes.NewBufferString("")
		req := httptest.NewRequest(http.MethodGet, "/users", reqBody)

		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		AssertStatus(t, res.Code, http.StatusOK)

	})
}

func NewMockServer() *http.ServeMux {
	mux := http.NewServeMux()

	userController := controllers.NewMockUserController()

	mux.HandleFunc("/users", userController.Index)

	return mux
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status got %d, want %d", got, want)
	}
}
