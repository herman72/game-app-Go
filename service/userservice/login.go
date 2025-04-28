package userservice

import (
	"fmt"
	"game-app-go/dto"
	"game-app-go/pkg/richerror"
)


func (s Service) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	// TODO - it would be better to user two separate method for existence check and getUserByPhoneNumber
	const op = "userservice.Login"
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return dto.LoginResponse{}, richerror.New(op).WithError(err).
		WithMeta(map[string]interface{}{"phone_number":req.PhoneNumber})
					
	}
	if !exist {
		return dto.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	if user.Password != getMD5Hash(req.Password) {
		return dto.LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	// TODO: generate random session, save session id in db, return sesion id to user
	accessToken, err := s.auth.CreateAccessToken(user)

	if err != nil{
		return dto.LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil{
		return dto.LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return dto.LoginResponse{
		User: dto.UserInfo{
			ID: user.ID,
			Name: user.Name,
			PhoneNumber: user.PhoneNumber,
		},
		Tokens: dto.Tokens{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}
