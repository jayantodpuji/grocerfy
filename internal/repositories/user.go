package repositories

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertRecord(context.Context, *models.User) error
	FindUserByEmail(context.Context, string) (*models.User, error)
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

func (u *userRepository) FindUserByEmail(c context.Context, email string) (*models.User, error) {
	var user models.User
	if err := u.db.WithContext(c).First(&user, "email = ?", email).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
