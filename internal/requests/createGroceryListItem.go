package requests

import "errors"

type CreateGroceryListItemRequest struct {
	GroceryListID string `json:"listId"`
	Name          string `json:"name"`
	Unit          string `json:"unit"`
	Quantity      int    `json:"quantity"`
	IsPurchased   bool   `json:"isPurchased"`
}

func (c *CreateGroceryListItemRequest) Validate() error {
	if c.GroceryListID == "" || c.Name == "" || c.Unit == "" || c.Quantity == 0 {
		return errors.New("invalid request")
	}

	return nil
}
