package userservice

import (
	"fmt"
	"game-app-go/entity"
	"game-app-go/pkg/phonenumber"
)

type Repository interface{
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Register(u entity.User)(entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type RegisterResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s Service) Register(req RegisterRequest)(RegisterResponse, error){
	// TODO: Verify phonenumber with verification code
	if !phonenumber.IsValid(req.PhoneNumber){
		return RegisterResponse{}, fmt.Errorf("phone number is not valid")
	}

	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected Error %w", err)
		}
	
		if !isUnique {
			return RegisterResponse{}, fmt.Errorf("phone number is not unique")
		}	
	}

	if len(req.Name) < 3 {
		return RegisterResponse{}, fmt.Errorf("length should be greater then 3")
	}
	user := entity.User{
		ID: 0,
		PhoneNumber: req.PhoneNumber,
		Name: req.Name,
	}
	createdUser, err := s.repo.Register(user)

	if err != nil{
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return RegisterResponse{
		User: createdUser,

	}, nil

	

}