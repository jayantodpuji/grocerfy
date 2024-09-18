package responses

import (
	"time"

	"github.com/gofrs/uuid"
)

type GroceryListItemDetail struct {
	ID            uuid.UUID `json:"id"`
	GroceryListID uuid.UUID `json:"groceryListID"`
	Category      string    `json:"category"`
	Name          string    `json:"name"`
	Unit          string    `json:"unit"`
	Size          int       `json:"size"`
	Quantity      int       `json:"quantity"`
	Price         float64   `json:"price"`
	PurchaseDate  time.Time `json:"purchaseDate"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}
