package responses

import (
	"time"

	"github.com/gofrs/uuid"
)

type GroceryListItemDetail struct {
	ID          uuid.UUID `json:"id"`
	Category    string    `json:"category"`
	Name        string    `json:"name"`
	Unit        string    `json:"unit"`
	Quantity    int       `json:"quantity"`
	Price       float64   `json:"price"`
	IsPurchased bool      `json:"isPurchased"`
	CreatedAt   time.Time `json:"createdAt"`
}
