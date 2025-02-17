package httpserver

import (
	"game-app-go/service/userservice"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server)userRegister(c echo.Context)error{
	

	var uReq userservice.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	Resp, err := s.userSvc.Register(uReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, Resp)
}