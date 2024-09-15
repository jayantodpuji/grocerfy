package handlers

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/labstack/echo/v4"
)

func GetUserIDFromContext(c echo.Context) (uuid.UUID, error) {
	userIDInterface := c.Get("user_id")
	if userIDInterface == nil {
		fmt.Println("Debug: userIDInterface is nil")
		return uuid.Nil, fmt.Errorf("user_id not found in context")
	}

	userIDString, ok := userIDInterface.(string)
	if !ok {
		return uuid.Nil, fmt.Errorf("invalid user_id format in context: %T", userIDInterface)
	}

	return uuid.FromString(userIDString)
}
