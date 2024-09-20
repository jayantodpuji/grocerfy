package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Email        string    `gorm:"type:text;unique;not null" json:"email"`
	PasswordHash string    `gorm:"type:text;not null" json:"passwordHash"`
	Name         string    `gorm:"type:text" json:"name"`
	CreatedAt    time.Time `gorm:"type:timestamptz;default:current_timestamp;not null" json:"createdAt"`
	UpdatedAt    time.Time `gorm:"type:timestamptz;default:current_timestamp;not null" json:"updatedAt"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.Must(uuid.NewV4())
	}
	return nil
}
