package authservice

import (
	"game-app-go/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Config struct {
	SignKey string `koanf:"sign_key"`
	AccessExpirationTime time.Duration `koanf:"access_expiration_time"`
	RefresExpirationTime time.Duration `koanf:"refresh_expiration_time"`
	AccessSubject string `koanf:"access_subject"`
	RefreshSubject string `koanf:"refresh_subject"`
}

type Service struct {
	config Config
}

func New(cfg Config)Service{

		return Service{
			config: cfg,
		}

}

func (s Service)CreateAccessToken(user entity.User)(string, error){

	return s.createToken(user.ID, s.config.AccessSubject, s.config.AccessExpirationTime)

}

func (s Service)CreateRefreshToken(user entity.User)(string, error){
	return s.createToken(user.ID, s.config.RefreshSubject, s.config.RefresExpirationTime)
	
}

func (s Service)ParseToken(bearerToken string)(* Claims, error){
	bearerToken = strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(bearerToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.config.SignKey), nil
	})

	if claims, ok := token.Claims.(*Claims); ok && token.Valid{
		return claims, nil
	} else {
		return nil, err
	}

	
}

func (s Service)createToken(userID uint, subject string, expireDuration time.Duration)(string, error){
	// TODO: replace with rca256

	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject: subject,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),

		},
		UserID: userID,
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString([]byte(s.config.SignKey)) 
	if err != nil{
		return "", err
	}
	return tokenString, nil
}