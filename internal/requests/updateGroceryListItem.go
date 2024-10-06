package requests

import "errors"

type UpdateGroceryListItem struct {
	Name        string `json:"name"`
	Unit        string `json:"unit"`
	Quantity    int    `json:"quantity"`
	IsPurchased bool   `json:"isPurchased"`
}

func (c *UpdateGroceryListItem) Validate() error {
	if c.Name == "" || c.Unit == "" || c.Quantity == 0 {
		return errors.New("invalid request")
	}

	return nil
}
