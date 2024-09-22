package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/delivery"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type GroceryListHandler interface {
	Create(echo.Context) error
	Index(echo.Context) error
	Detail(echo.Context) error
	Update(echo.Context) error
	Delete(echo.Context) error
}

type groceryListHandler struct {
	groceryListService services.GroceryListService
}

type GroceryListHandlerDependency struct {
	GroceryListService services.GroceryListService
}

func NewGroceryListHandler(deps GroceryListHandlerDependency) GroceryListHandler {
	return &groceryListHandler{groceryListService: deps.GroceryListService}
}

func (g *groceryListHandler) Create(c echo.Context) error {
	var req requests.CreateGroceryListRequest
	if err := c.Bind(&req); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	if err := req.Validate(); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return delivery.ResponseError(c, http.StatusUnauthorized, "invalid or missing user authentication")
	}

	if err := g.groceryListService.CreateGroceryList(c.Request().Context(), userID, req); err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

func (g *groceryListHandler) Index(c echo.Context) error {
	userID, err := GetUserIDFromContext(c)
	if err != nil {
		return delivery.ResponseError(c, http.StatusUnauthorized, "invalid or missing user authentication")
	}

	groceryLists, err := g.groceryListService.GetGroceryListByUserID(c.Request().Context(), userID)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return delivery.ResponseSuccess(c, http.StatusOK, groceryLists)
}

func (g *groceryListHandler) Detail(c echo.Context) error {
	id := c.Param("id")

	listID, err := uuid.FromString(id)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	groceryList, err := g.groceryListService.GetGroceryListByID(c.Request().Context(), listID)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return delivery.ResponseSuccess(c, http.StatusOK, groceryList)
}

func (g *groceryListHandler) Update(c echo.Context) error {
	idAny := c.Param("id")

	id, err := uuid.FromString(idAny)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	var req requests.UpdateGroceryListRequest
	if err := c.Bind(&req); err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = g.groceryListService.UpdateGroceryList(c.Request().Context(), id, &req)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (g *groceryListHandler) Delete(c echo.Context) error {
	idAny := c.Param("id")

	id, err := uuid.FromString(idAny)
	if err != nil {
		return delivery.ResponseError(c, http.StatusBadRequest, err.Error())
	}

	err = g.groceryListService.DestroyGroceryList(c.Request().Context(), id)
	if err != nil {
		return delivery.ResponseError(c, http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}
