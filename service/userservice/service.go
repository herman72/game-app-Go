package userservice

import (
	"crypto/md5"
	"encoding/hex"
	"game-app-go/entity"
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

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

