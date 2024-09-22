package requests

import "errors"

type UpdateGroceryListItem struct {
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Unit        string  `json:"unit"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	IsPurchased bool    `json:"isPurchased"`
}

// TODO
// need to validate correctly only if key is sent from FE
func (c *UpdateGroceryListItem) Validate() error {
	if c.Category == "" || c.Name == "" || c.Unit == "" || c.Quantity == 0 || c.Price == 0 {
		return errors.New("invalid request")
	}

	return nil
}
