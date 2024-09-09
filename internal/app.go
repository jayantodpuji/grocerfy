package internal

import (
	"context"

	"github.com/jayantodpuji/grocerfy/config"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
}

type Application struct {
	DB     *gorm.DB
	Router *echo.Echo
	Port   string
	Env    string
	Handler
}

func NewApplication(db *gorm.DB, ec *echo.Echo, cfg *config.Config) (*Application, error) {
	return &Application{
		Handler: Handler{},
		Router:  ec,
		Port:    cfg.App.Port,
		Env:     cfg.App.Env,
	}, nil
}

func (a *Application) Start() error {
	return a.Router.Start(a.Port)
}

func (a *Application) Shutdown(c context.Context) error {
	return a.Router.Shutdown(c)
}
