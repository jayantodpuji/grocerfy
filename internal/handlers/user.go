package handlers

import (
	"net/http"

	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	SignUp(echo.Context) error
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

func (uh *userHandler) SignUp(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
