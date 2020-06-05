package controllers

import (
	"encoding/json"
	"net/http"

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
	if err != err {
		w.Header().Set("Content-Type", "applicaiton/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
