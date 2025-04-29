package userhandler

import (
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"
	"game-app-go/validator/uservalidator"

)

type Handler struct{
	authConfig authservice.Config
	authSvc authservice.Service
	userSvc userservice.Service
	userValidator uservalidator.Validator

}

func New(authSvc authservice.Service, userSvc userservice.Service, 
	userValidator uservalidator.Validator,
	authConfig authservice.Config) Handler {
	return Handler{
		authSvc: authSvc,
		userSvc: userSvc,
		userValidator: userValidator,
		authConfig: authConfig,
	}
}




