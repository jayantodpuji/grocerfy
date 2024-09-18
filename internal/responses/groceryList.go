package responses

import "time"

type GroceryListIndexResponse struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

type GroceryListDetailResponse struct {
	ID          string                  `json:"id"`
	Name        string                  `json:"name"`
	Description string                  `json:"description"`
	CreatedAt   time.Time               `json:"createdAt"`
	Items       []GroceryListItemDetail `json:"items,omitempty"`
}
