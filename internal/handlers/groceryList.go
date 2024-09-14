package handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	"github.com/labstack/echo/v4"
)

type GroceryListHandler interface {
	Create(c echo.Context) error
	Index(c echo.Context) error
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
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if err := req.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	if err := g.groceryListService.CreateGroceryList(c.Request().Context(), req); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "Grocery list created successfully",
	})
}

func (g *groceryListHandler) Index(c echo.Context) error {
	userIDString := c.Get("user_id").(*jwt.RegisteredClaims).Subject
	if userIDString == "" {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Unauthorized",
		})
	}

	userID, err := uuid.FromString(userIDString)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": err.Error(),
		})
	}

	groceryLists, err := g.groceryListService.GetGroceryListByUserID(c.Request().Context(), userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": groceryLists,
	})
}
