package infrastructure

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/go-chi/chi"
	"github.com/hiroshimashu/ei-rest/app/domain"
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

	t.Run("Successfully post user", func(t *testing.T) {
		reqBody := bytes.NewBufferString(`{"ID":"7777","Gender":"M","Age":26,"Area":"tokyo","Email":"example3@g.com","Password":"1111","IsPaid":true}`)
		req := httptest.NewRequest(http.MethodPost, "/user", reqBody)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		got := getUserFromResponse(t, res.Body)
		want := domain.User{
			ID:       "7777",
			Gender:   "M",
			Age:      26,
			Area:     "tokyo",
			Email:    "example3@g.com",
			Password: "1111",
			IsPaid:   true,
		}
		assertUser(t, got, want)

	})
}

func TestMovieRouter(t *testing.T) {
	server := NewMockServer()

	t.Run("it returns 200 on /movies", func(t *testing.T) {
		reqBody := bytes.NewBufferString("")
		req := httptest.NewRequest(http.MethodGet, "/movies", reqBody)

		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		AssertStatus(t, res.Code, http.StatusOK)

	})

	t.Run("it returns 200 on /movie/{id}", func(t *testing.T) {
		reqBody := bytes.NewBufferString("")
		req := httptest.NewRequest(http.MethodGet, "/movie/5555", reqBody)
		res := httptest.NewRecorder()
		server.ServeHTTP(res, req)
		AssertStatus(t, res.Code, http.StatusOK)
	})

}

func NewMockServer() *chi.Mux {
	r := chi.NewRouter()

	userController := controllers.NewMockUserController()
	movieController := controllers.NewMockMovieController()

	r.Get("/users", userController.Index)
	r.Post("/user", userController.Create)
	r.Get("/movies", movieController.Index)
	r.Route("/movie", func(r chi.Router) {
		r.Post("/", movieController.Create)
		r.Route("/{movieID}", func(r chi.Router) {
			r.Get("/", movieController.IndexByID)
		})
	})

	return r
}

func AssertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status got %d, want %d", got, want)
	}
}

func getUserFromResponse(t *testing.T, body io.Reader) (user domain.User) {
	t.Helper()

	err := json.NewDecoder(body).Decode(&user)

	if err != nil {
		t.Fatal("unable to parse response from server")
	}

	return
}

func assertUser(t *testing.T, got, want domain.User) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
