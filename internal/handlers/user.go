package handlers

import (
	"net/http"

	"github.com/jayantodpuji/grocerfy/internal/delivery"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	Register(echo.Context) error
	Login(echo.Context) error
}

type userHandler struct {
	userService services.UserService
}

type UserHandlerDependency struct {
	UserService services.UserService
}

func NewUserHandler(deps UserHandlerDependency) UserHandler {
	return &userHandler{userService: deps.UserService}
}

func (uh *userHandler) Register(c echo.Context) error {
	var req requests.UserRegistration
	err := c.Bind(&req)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = req.Validate()
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = uh.userService.Register(c.Request().Context(), req)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (uh *userHandler) Login(c echo.Context) error {
	var req requests.UserLogin
	err := c.Bind(&req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	token, err := uh.userService.Login(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
	})
}
