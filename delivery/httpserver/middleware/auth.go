package middleware

import (
	"game-app-go/service/authservice"
	cfg "game-app-go/config"

	mv "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// clouser or higher order function
func Auth(service authservice.Service, config authservice.Config) echo.MiddlewareFunc {
	return mv.WithConfig(mv.Config{
		ContextKey: cfg.AuthMiddlewareContextKey,
		SigningKey: []byte(config.SignKey),
		SigningMethod: "HS256",
		// TODO: as sign method to config
		ParseTokenFunc: func(c echo.Context, auth string) (interface{}, error) {
			claims, err := service.ParseToken(auth)
			if err != nil {
				return nil, err
			}
			return claims, nil
		},
	})
}