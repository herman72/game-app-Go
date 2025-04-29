package userhandler

import (
	"game-app-go/param"
	"game-app-go/pkg/errmsg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)


func (h Handler) userLogin(c echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}

	if err, fieldErrors   := h.userValidator.ValidateLoginRequest(req); err != nil {

		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := h.userSvc.Login(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}

	return c.JSON(http.StatusOK, resp)

}
