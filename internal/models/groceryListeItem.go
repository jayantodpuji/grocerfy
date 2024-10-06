package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

// TODO
// How to handle price decimal
type GroceryListItem struct {
	ID            uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	GroceryListID uuid.UUID `gorm:"type:uuid;not null"`
	Name          string    `gorm:"type:varchar(255);not null"`
	Unit          string    `gorm:"type:varchar(255);not null"`
	Quantity      int       `gorm:"type:int;not null"`
	IsPurchased   bool      `gorm:"type:boolean;not null"`
	CreatedAt     time.Time `gorm:"type:timestamptz;not null"`
	UpdatedAt     time.Time `gorm:"type:timestamptz;not null"`
}

func (g *GroceryListItem) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.Must(uuid.NewV4())
	}
	return nil
}
