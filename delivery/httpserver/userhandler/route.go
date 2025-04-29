package userhandler

import (
	"game-app-go/delivery/httpserver/middleware"

	"github.com/labstack/echo/v4"
)

func (h Handler) SetUserRoutes(e *echo.Echo) {
	userGroup := e.Group("/users")

	
	userGroup.POST("/login", h.userLogin, middleware.Auth(h.authSvc, h.authConfig))
	userGroup.POST("/register", h.userRegister, middleware.Auth(h.authSvc, h.authConfig))
	userGroup.GET("/profile", h.userProfile, middleware.Auth(h.authSvc, h.authConfig))
}