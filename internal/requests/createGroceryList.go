package requests

import (
	"errors"

	"github.com/gofrs/uuid"
)

type CreateGroceryListRequest struct {
	Name        string    `json:"name" validate:"required"`
	UserID      uuid.UUID `json:"userID" validate:"required"`
	Description string    `json:"description"`
}

func (c *CreateGroceryListRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	if c.UserID == uuid.Nil {
		return errors.New("userID is required")
	}

	return nil
}
