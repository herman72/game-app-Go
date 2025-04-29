package uservalidator

import "game-app-go/entity"


const (
	phoneNumberRegex = "^09[0-9]{9}$"
)

type Repository interface{
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, error)
}

type Validator struct {
	repo Repository
}

func New(repo Repository)Validator{
	return Validator{repo: repo}
}


