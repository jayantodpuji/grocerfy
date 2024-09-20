package requests

import (
	"errors"
)

type CreateGroceryListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *CreateGroceryListRequest) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
