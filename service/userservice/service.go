package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"game-app-go/dto"
	"game-app-go/entity"
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

func New(authGenerator AuthGenerator, repo Repository) Service {
	return Service{auth: authGenerator ,repo: repo} 
}

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

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string	`json:"password"`
}

type Tokens struct {
	AccessToken string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	User dto.UserInfo `json:"user"`
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
		User: dto.UserInfo{
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



