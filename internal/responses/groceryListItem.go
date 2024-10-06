package responses

import (
	"time"

	"github.com/gofrs/uuid"
)

type GroceryListItemDetail struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Unit        string    `json:"unit"`
	Quantity    int       `json:"quantity"`
	IsPurchased bool      `json:"isPurchased"`
	CreatedAt   time.Time `json:"createdAt"`
}
