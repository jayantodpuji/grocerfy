package handlers

import (
	"net/http"

	"github.com/jayantodpuji/grocerfy/internal/delivery"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type GroceryListItemHandler interface {
	Create(echo.Context) error
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

func (h *groceryListItemHandler) Create(c echo.Context) error {
	var req requests.CreateGroceryListItemRequest

	if err := c.Bind(&req); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	if err := h.service.CreateGroceryListItem(c.Request().Context(), &req); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	return delivery.ResponseSuccess(c, http.StatusOK, nil)
}
