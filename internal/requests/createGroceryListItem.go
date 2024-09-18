package requests

import "errors"

type CreateGroceryListItemRequest struct {
	GroceryListID string  `json:"groceryListID"`
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	Unit          string  `json:"unit"`
	Size          int     `json:"size"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
}

// TODO:
// Need to show error on specific field missing
func (c *CreateGroceryListItemRequest) Validate() error {
	if c.GroceryListID == "" || c.Category == "" || c.Name == "" || c.Unit == "" || c.Size == 0 || c.Quantity == 0 || c.Price == 0 {
		return errors.New("invalid request")
	}

	return nil
}
