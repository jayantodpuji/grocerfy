package handlers

import (
	"net/http"

	"github.com/jayantodpuji/grocerfy/internal/delivery"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type GroceryListItemHandler interface {
	GetByGroceryList(echo.Context) error
}

type groceryListItemHandler struct {
	service services.GroceryListItemService
}

type GroceryListItemHandlerDependency struct {
	Service services.GroceryListItemService
}

func NewGroceryListItemHandler(deps GroceryListItemHandlerDependency) GroceryListItemHandler {
	return &groceryListItemHandler{service: deps.Service}
}

func (h *groceryListItemHandler) GetByGroceryList(c echo.Context) error {
	return delivery.ResponseSuccess(c, http.StatusOK, nil)
}
