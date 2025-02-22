package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game-app-go/entity"
	"game-app-go/pkg/phonenumber"
	"game-app-go/pkg/richerror"
)

type Repository interface {
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(PhoneNumber string) (entity.User, bool, error)
	GetUserByID(userID uint)(entity.User, error)
}

type AuthGenerator interface{
	CreateAccessToken(user entity.User)(string, error)
	CreateRefreshToken(user entity.User)(string, error)
}

type Service struct {
	auth AuthGenerator
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type UserInfo struct {
	ID uint `json:"id"`
	PhoneNumber string `json:"phone_number"`
	Name string `json:"name"`
}
type RegisterResponse struct {
	User UserInfo `json:"user"`
			
}

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator ,repo: repo} 
}

func (s Service) Register(req RegisterRequest) (RegisterResponse, error) {
	// TODO: Verify phonenumber with verification code
	if !phonenumber.IsValid(req.PhoneNumber) {
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
	if len(req.Password) < 8 {
		return RegisterResponse{}, fmt.Errorf("password length should be greater then 8")
	}

	// TODO: use bcycpt for hashing

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    getMD5Hash(req.Password),
	}
	createdUser, err := s.repo.Register(user)

	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return RegisterResponse{UserInfo{
		ID: createdUser.ID,
		PhoneNumber: createdUser.PhoneNumber,
		Name: createdUser.Name,
	}}, nil

}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string	`json:"password"`
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	User UserInfo `json:"user"`
	Tokens Tokens `json:"tokens"`
	

}

func (s Service) Login(req LoginRequest) (LoginResponse, error) {
	// TODO - it would be better to user two separate method for existence check and getUserByPhoneNumber
	const op = "userservice.Login"
	user, exist, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, richerror.New(op).WithError(err).
		WithMeta(map[string]interface{}{"phone_number":req.PhoneNumber})
					
	}
	if !exist {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	if user.Password != getMD5Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password isn't correct")
	}
	// TODO: generate random session, save session id in db, return sesion id to user
	accessToken, err := s.auth.CreateAccessToken(user)

	if err != nil{
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	refreshToken, err := s.auth.CreateRefreshToken(user)
	if err != nil{
		return LoginResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return LoginResponse{
		User: UserInfo{
			ID: user.ID,
			Name: user.Name,
			PhoneNumber: user.PhoneNumber,
		},
		Tokens: Tokens{
			AccessToken: accessToken,
			RefreshToken: refreshToken,
		},
	}, nil
}

type ProfileRequest struct {

	UserID uint

}

type ProfileResponse struct {
	Name string `json:"name"`
}

// All req inputs for intractor/service should be sanitized

func (s Service)Profile(req ProfileRequest)(ProfileResponse, error){
	const op =  "userservice.Profile"
	user, err := s.repo.GetUserByID(req.UserID)
	// I assume data is already sanitized
	if err != nil{
		// TODO: we can use rich error for better error handeling
		return ProfileResponse{}, richerror.New(op).
											WithError(err).
											WithMeta(map[string]interface{}{"req":req})
		
		// return ProfileResponse{}, fmt.Errorf("unexpected error %w", err)
	}

	return ProfileResponse{Name: user.Name}, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}



