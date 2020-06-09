package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/hiroshimashu/ei-rest/app/domain"
	"github.com/hiroshimashu/ei-rest/app/interfaces/database"
	"github.com/hiroshimashu/ei-rest/app/usecases"
)

type MockUserController struct {
	Interactor usecases.UserInteractor
}

func NewMockUserController() *MockUserController {
	return &MockUserController{
		Interactor: usecases.UserInteractor{
			UserRepository: database.NewImMemoryUserRepositry(),
		},
	}
}

func (mc *MockUserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := mc.Interactor.Index()
	if err != nil {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (mc *MockUserController) Create(w http.ResponseWriter, r *http.Request) {
	var u domain.User
	err := json.NewDecoder(r.Body).Decode(&u)
	user, err := mc.Interactor.Add(u)
	if err != nil {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
