package controllers

import (
	"github.com/hiroshimashu/ei-rest/app/usecases"
)

type UserController struct {
	Interactor usecases.UserInteractor
}
