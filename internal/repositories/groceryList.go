package repositories

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"gorm.io/gorm"
)

type GroceryListRepository interface {
	InsertRecord(context.Context, *models.GroceryList) error
	GetGroceryListByUserID(context.Context, uuid.UUID) ([]models.GroceryList, error)
}

type groceryListRepository struct {
	db *gorm.DB
}

type GroceryListRepositoryDependency struct {
	DB *gorm.DB
}

func NewGroceryListRepository(deps GroceryListRepositoryDependency) GroceryListRepository {
	return &groceryListRepository{db: deps.DB}
}

func (g *groceryListRepository) InsertRecord(c context.Context, p *models.GroceryList) error {
	if err := g.db.WithContext(c).Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (g *groceryListRepository) GetGroceryListByUserID(c context.Context, userID uuid.UUID) ([]models.GroceryList, error) {
	var groceryLists []models.GroceryList
	if err := g.db.WithContext(c).Where("user_id = ?", userID).Find(&groceryLists).Error; err != nil {
		return nil, err
	}

	return groceryLists, nil
}
