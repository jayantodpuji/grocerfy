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
	UpdateItemDetail(context.Context, uuid.UUID, *requests.UpdateGroceryListItem) error
	DeleteItem(context.Context, uuid.UUID) error
	ToggleIsPurchased(context.Context, uuid.UUID) error
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
		Name:          req.Name,
		Unit:          req.Unit,
		Quantity:      req.Quantity,
	})

	return err
}

func (s *groceryListItemService) UpdateItemDetail(c context.Context, id uuid.UUID, req *requests.UpdateGroceryListItem) error {
	p := models.GroceryListItem{
		Name:        req.Name,
		Unit:        req.Unit,
		Quantity:    req.Quantity,
		IsPurchased: req.IsPurchased,
	}

	err := s.repo.UpdateItemByID(c, id, p)
	if err != nil {
		return err
	}

	return nil
}

func (s groceryListItemService) DeleteItem(c context.Context, id uuid.UUID) error {
	if err := s.repo.DestroyItemByID(c, id); err != nil {
		return err
	}

	return nil
}

func (s groceryListItemService) ToggleIsPurchased(c context.Context, id uuid.UUID) error {
	return s.repo.ToggleIsPurchased(c, id)
}
