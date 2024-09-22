package repositories

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"gorm.io/gorm"
)

type GroceryListItemRepository interface {
	InsertRecord(context.Context, *models.GroceryListItem) error
	GetItemsByGroceryList(context.Context, uuid.UUID) ([]*models.GroceryListItem, error)
	GetItemByID(context.Context, uuid.UUID) (*models.GroceryListItem, error)
	UpdateItemByID(context.Context, uuid.UUID, models.GroceryListItem) error
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

func (g *groceryListItemRepository) InsertRecord(c context.Context, p *models.GroceryListItem) error {
	if err := g.db.WithContext(c).Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func (g *groceryListItemRepository) GetItemsByGroceryList(c context.Context, listID uuid.UUID) ([]*models.GroceryListItem, error) {
	var items []*models.GroceryListItem
	if err := g.db.WithContext(c).Where("grocery_list_id = ?", listID.String()).Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}

func (g *groceryListItemRepository) GetItemByID(c context.Context, id uuid.UUID) (*models.GroceryListItem, error) {
	var item models.GroceryListItem
	if err := g.db.WithContext(c).First(&item).Error; err != nil {
		return nil, err
	}

	return &item, nil
}

func (g *groceryListItemRepository) UpdateItemByID(c context.Context, id uuid.UUID, p models.GroceryListItem) error {
	if err := g.db.WithContext(c).
		Model(&models.GroceryListItem{}).
		Where("id = ?", id).
		Updates(p).Error; err != nil {
		return err
	}

	return nil
}
