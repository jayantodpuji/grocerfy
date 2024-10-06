package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/delivery"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type GroceryListItemHandler interface {
	Create(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
	Toggle(echo.Context) error
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

	return c.NoContent(http.StatusOK)
}

func (h *groceryListItemHandler) Update(c echo.Context) error {
	idAny := c.Param("id")

	id, err := uuid.FromString(idAny)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	var req requests.UpdateGroceryListItem
	if err := c.Bind(&req); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = h.service.UpdateItemDetail(c.Request().Context(), id, &req)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *groceryListItemHandler) Toggle(c echo.Context) error {
	idAny := c.Param("id")

	id, err := uuid.FromString(idAny)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = h.service.ToggleIsPurchased(c.Request().Context(), id)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *groceryListItemHandler) Delete(c echo.Context) error {
	idAny := c.Param("id")

	id, err := uuid.FromString(idAny)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = h.service.DeleteItem(c.Request().Context(), id)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
