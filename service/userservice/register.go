package userservice

import (
	"fmt"
	"game-app-go/param"
	"game-app-go/entity"
)

func (s Service) Register(req dto.RegisterRequest) (dto.RegisterResponse, error) {
	// TODO: Verify phonenumber with verification code

	// TODO: use bcycpt for hashing

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}
	createdUser, err := s.repo.Register(user)

	if err != nil {
		return dto.RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return dto.RegisterResponse{User: dto.UserInfo{
		ID: createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name: createdUser.Name,
	}}, nil

}
