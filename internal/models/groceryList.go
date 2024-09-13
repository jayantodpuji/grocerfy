package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type GroceryList struct {
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID      uuid.UUID  `gorm:"not null" json:"userId"`
	User        User       `gorm:"foreignKey:UserID" json:"user"`
	Name        string     `gorm:"type:text;not null" json:"name"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;not null" json:"createdAt"`
	UpdatedAt   time.Time  `gorm:"type:timestamptz;not null" json:"updatedAt"`
	DeletedAt   *time.Time `gorm:"type:timestamptz" json:"deletedAt"`
}

func (g *GroceryList) BeforeCreate(tx *gorm.DB) error {
	if g.ID == uuid.Nil {
		g.ID = uuid.Must(uuid.NewV4())
	}
	return nil
}
