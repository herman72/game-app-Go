package httpserver

import (
	"fmt"
	"game-app-go/config"
	"game-app-go/service/authservice"
	"game-app-go/service/userservice"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	config config.Config
	authSvc authservice.Service
	userSvc userservice.Service
}

func New(config config.Config, authSvc authservice.Service, userSvc userservice.Service)Server{
	return Server{
		config: config,
		authSvc: authSvc,
		userSvc: userSvc,
	}
	}


func (s Server)Serve(){
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userGroup := e.Group("/users")

	e.GET("/health-check", s.healthCheck)
	userGroup.POST("/login", s.userLogin)
	userGroup.POST("/register", s.userRegister)
	userGroup.GET("/profile", s.userProfile)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", s.config.HTTPServer.Port)))

}