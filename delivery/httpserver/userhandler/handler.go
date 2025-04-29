package userhandler

import (
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"
	"game-app-go/validator/uservalidator"

)

type Handler struct{
	authSvc authservice.Service
	userSvc userservice.Service
	userValidator uservalidator.Validator

}

func New(authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator) Handler {
	return Handler{
		authSvc: authSvc,
		userSvc: userSvc,
		userValidator: userValidator,
	}
}




