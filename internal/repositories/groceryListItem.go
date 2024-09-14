package repositories

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"gorm.io/gorm"
)

type GroceryListItemRepository interface {
	InsertRecord(context.Context, *models.GroceryListItem) (*models.GroceryListItem, error)
	InsertRecords(context.Context, []*models.GroceryListItem) ([]*models.GroceryListItem, error)
}

type groceryListItemRepository struct {
	db *gorm.DB
}

type GroceryListItemRepositoryDependency struct {
	DB *gorm.DB
}

func NewGroceryListItemRepository(deps GroceryListItemRepositoryDependency) GroceryListItemRepository {
	return &groceryListItemRepository{db: deps.DB}
}

func (g *groceryListItemRepository) InsertRecord(c context.Context, p *models.GroceryListItem) (*models.GroceryListItem, error) {
	if err := g.db.WithContext(c).Create(&p).Error; err != nil {
		return nil, err
	}

	return p, nil
}

func (g *groceryListItemRepository) InsertRecords(c context.Context, p []*models.GroceryListItem) ([]*models.GroceryListItem, error) {
	err := g.db.WithContext(c).Transaction(func(tx *gorm.DB) error {
		return tx.CreateInBatches(&p, 10).Error
	})

	if err != nil {
		return nil, err
	}

	return p, nil
}
