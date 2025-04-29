package httpserver

import (
	"fmt"
	"game-app-go/config"
	"game-app-go/delivery/httpserver/userhandler"
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"
	"game-app-go/validator/uservalidator"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	userhandler userhandler.Handler
	
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service, userValidator uservalidator.Validator)Server{
	return Server{
		config: config,
		userhandler: userhandler.New(authSvc, userSvc, userValidator, config.Auth),
		
	}
	}


func (s Server)Serve(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/health-check", s.healthCheck)

	s.userhandler.SetUserRoutes(e)

	

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}