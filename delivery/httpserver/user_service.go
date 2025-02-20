package httpserver

import (
	"game-app-go/pkg/errmsg/httpmsg"
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

func (s Server)userLogin(c echo.Context) error {

	
	var req userservice.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	
	resp, err := s.userSvc.Login(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		
	}

	return c.JSON(http.StatusOK, resp)

}

func (s Server)userProfile(c echo.Context)error{
	// 	// sessionID := req.Header.Get("SessionID")
// 	// TODO: Validate sessionid by database and get user id

	// validate jwt token and retrive userID from pyload

	
	authToken := c.Request().Header.Get("Authorization")	
	claims, err := s.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	

	resp, err := s.userSvc.Profile(userservice.ProfileRequest{UserID: claims.UserID})

	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)


}