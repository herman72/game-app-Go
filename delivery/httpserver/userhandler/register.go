package userhandler

import (
	"game-app-go/param"
	"game-app-go/pkg/errmsg/httpmsg"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) userRegister(c echo.Context) error {

	var uReq dto.RegisterRequest

	if err := c.Bind(&uReq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErrors, err  := h.userValidator.ValidateRegisterRequest(uReq); err != nil {

		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	Resp, err := h.userSvc.Register(uReq)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, Resp)
}