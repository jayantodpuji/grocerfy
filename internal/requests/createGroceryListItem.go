package requests

import "errors"

// TODO:
// - all request key type should be string then parse it to right type
type CreateGroceryListItemRequest struct {
	GroceryListID string  `json:"listId"`
	Category      string  `json:"category"`
	Name          string  `json:"name"`
	Unit          string  `json:"unit"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	IsPurchased   bool    `json:"isPurchased"`
}

// TODO:
// Need to show error on specific field missing
func (c *CreateGroceryListItemRequest) Validate() error {
	if c.GroceryListID == "" || c.Name == "" || c.Unit == "" || c.Quantity == 0 {
		return errors.New("invalid request")
	}

	return nil
}
