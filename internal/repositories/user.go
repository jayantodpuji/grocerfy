package repositories

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertRecord(context.Context, *models.User) error
}

type userRepository struct {
	db *gorm.DB
}

type UserRepositoryDependency struct {
	DB *gorm.DB
}

func NewUserRepository(deps UserRepositoryDependency) UserRepository {
	return &userRepository{db: deps.DB}
}

func (u *userRepository) InsertRecord(c context.Context, p *models.User) error {
	if err := u.db.WithContext(c).Create(&p).Error; err != nil {
		return err
	}

	return nil
}
