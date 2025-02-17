package authservice

import (
	"game-app-go/entity"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	signKey string
	accessExpirationTime time.Duration
	refresExpirationTime time.Duration
	accessSubject string
	refreshSubject string
}

func New(signKey, accessSubject, refreshSubject string, 
	accessExpirationTime, refresExpirationTime time.Duration)Service{

		return Service{
			signKey: signKey,
			accessExpirationTime: accessExpirationTime,
			refresExpirationTime: refresExpirationTime,
			accessSubject: accessSubject,
			refreshSubject: refreshSubject,
		}

}

func (s Service)CreateAccessToken(user entity.User)(string, error){

	return s.createToken(user.ID, s.accessSubject, s.accessExpirationTime)

}

func (s Service)CreateRefreshToken(user entity.User)(string, error){
	return s.createToken(user.ID, s.refreshSubject, s.refresExpirationTime)
	
}

func (s Service)ParseToken(bearerToken string)(* Claims, error){
	bearerToken = strings.Replace(bearerToken, "Bearer ", "", 1)
	token, err := jwt.ParseWithClaims(bearerToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.signKey), nil
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
	tokenString, err := accessToken.SignedString([]byte(s.signKey)) 
	if err != nil{
		return "", err
	}
	return tokenString, nil
}