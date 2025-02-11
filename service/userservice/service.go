package userservice

import (
	"crypto/md5"
	"encoding/hex"
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
	Password string `json:"password"`
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

	// TODO: validate password - use regex
	if len(req.Password)<8{
		return RegisterResponse{}, fmt.Errorf("password length should be greater then 8")
	}
	
	// TODO: use bcycpt for hashing
	
	user := entity.User{
		ID: 0,
		PhoneNumber: req.PhoneNumber,
		Name: req.Name,
		Password: getMD5Hash(req.Password),
	}
	createdUser, err := s.repo.Register(user)

	if err != nil{
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return RegisterResponse{
		User: createdUser,

	}, nil

	

}

type LoginRequest struct {
	PhoneNumber string
	Password string
}

type LoginResponse struct {

}

func(s Service)Login(req LoginRequest)(LoginResponse, error){
	panic("implement me")
}

func getMD5Hash(text string)string{
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}