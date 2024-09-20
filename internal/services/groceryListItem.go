package services

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
)

type GroceryListItemService interface {
	CreateGroceryListItem(context.Context, *requests.CreateGroceryListItemRequest) error
}

type groceryListItemService struct {
	repo repositories.GroceryListItemRepository
}

func NewGroceryListItemService(repo repositories.GroceryListItemRepository) GroceryListItemService {
	return &groceryListItemService{repo: repo}
}

func (s *groceryListItemService) CreateGroceryListItem(c context.Context, req *requests.CreateGroceryListItemRequest) error {
	err := s.repo.InsertRecord(c, &models.GroceryListItem{
		GroceryListID: uuid.Must(uuid.FromString(req.GroceryListID)),
		Category:      req.Category,
		Name:          req.Name,
		Unit:          req.Unit,
		Quantity:      req.Quantity,
		Price:         req.Price,
	})

	return err
}
