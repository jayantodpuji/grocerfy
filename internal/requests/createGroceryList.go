package requests

import (
	"errors"

	"github.com/gofrs/uuid"
)

type CreateGroceryListRequest struct {
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"userID"`
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
