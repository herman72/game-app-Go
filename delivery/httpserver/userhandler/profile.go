package userhandler

import (
	"game-app-go/param"
	"game-app-go/pkg/errmsg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userProfile(c echo.Context) error {
	// 	// sessionID := req.Header.Get("SessionID")
	// 	// TODO: Validate sessionid by database and get user id

	// validate jwt token and retrive userID from pyload

	authToken := c.Request().Header.Get("Authorization")
	claims, err := h.authSvc.ParseToken(authToken)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}

	resp, err := h.userSvc.Profile(dto.ProfileRequest{UserID: claims.UserID})

	if err != nil {
		msg, code := httpmsg.Error(err)
		return echo.NewHTTPError(code, msg)
	}

	return c.JSON(http.StatusOK, resp)

}