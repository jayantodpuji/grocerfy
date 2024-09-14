package services

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
)

type GroceryListService interface {
	CreateGroceryList(context.Context, requests.CreateGroceryListRequest) error
	GetGroceryListByUserID(context.Context, uuid.UUID) ([]models.GroceryList, error)
}

type groceryListService struct {
	groceryListRepository repositories.GroceryListRepository
}

type GroceryListServiceDependency struct {
	GroceryListRepository repositories.GroceryListRepository
}

func NewGroceryListService(deps GroceryListServiceDependency) GroceryListService {
	return &groceryListService{groceryListRepository: deps.GroceryListRepository}
}

func (g *groceryListService) CreateGroceryList(c context.Context, p requests.CreateGroceryListRequest) error {
	if err := g.groceryListRepository.InsertRecord(c, &models.GroceryList{
		Name:        p.Name,
		UserID:      p.UserID,
		Description: p.Description,
	}); err != nil {
		return err
	}

	return nil
}

func (g *groceryListService) GetGroceryListByUserID(c context.Context, userID uuid.UUID) ([]models.GroceryList, error) {
	return g.groceryListRepository.GetGroceryListByUserID(c, userID)
}
