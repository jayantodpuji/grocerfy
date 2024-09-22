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
	GetGroceryListByID(context.Context, uuid.UUID) (*models.GroceryList, error)
	UpdateGroceryListByID(context.Context, uuid.UUID, *models.GroceryList) error
	DestroyGroceryListAndItemsByID(context.Context, uuid.UUID) error
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

func (g *groceryListRepository) GetGroceryListByID(c context.Context, listID uuid.UUID) (*models.GroceryList, error) {
	var groceryList models.GroceryList
	if err := g.db.WithContext(c).Where("id = ?", listID).First(&groceryList).Error; err != nil {
		return nil, err
	}

	return &groceryList, nil
}

func (g *groceryListRepository) UpdateGroceryListByID(c context.Context, id uuid.UUID, p *models.GroceryList) error {
	if err := g.db.WithContext(c).
		Model(&models.GroceryList{}).
		Where("id = ?", id).
		Updates(p).Error; err != nil {
		return err
	}

	return nil
}

func (g *groceryListRepository) DestroyGroceryListAndItemsByID(c context.Context, groceryListID uuid.UUID) error {
	tx := g.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.WithContext(c).
		Unscoped().
		Where("grocery_list_id = ?", groceryListID).
		Delete(&models.GroceryListItem{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.WithContext(c).
		Unscoped().
		Delete(&models.GroceryList{}, groceryListID).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
