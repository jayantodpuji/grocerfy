package services

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/responses"
)

type GroceryListItemService interface {
	CreateGroceryListItem(context.Context, *requests.CreateGroceryListItemRequest) error
	GetGroceryListItemDetail(context.Context, uuid.UUID) (*responses.GroceryListItemDetail, error)
	UpdateItemDetail(context.Context, uuid.UUID, *requests.UpdateGroceryListItem) error
	DeleteItem(context.Context, uuid.UUID) error
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

func (s *groceryListItemService) GetGroceryListItemDetail(c context.Context, id uuid.UUID) (*responses.GroceryListItemDetail, error) {
	item, err := s.repo.GetItemByID(c, id)
	if err != nil {
		return nil, err
	}

	return &responses.GroceryListItemDetail{
		ID:          item.ID,
		Category:    item.Category,
		Name:        item.Name,
		Unit:        item.Unit,
		Quantity:    item.Quantity,
		Price:       item.Price,
		IsPurchased: item.IsPurchased,
		CreatedAt:   item.CreatedAt,
	}, nil
}

func (s *groceryListItemService) UpdateItemDetail(c context.Context, id uuid.UUID, req *requests.UpdateGroceryListItem) error {
	p := models.GroceryListItem{
		Category:    req.Category,
		Name:        req.Name,
		Unit:        req.Unit,
		Quantity:    req.Quantity,
		Price:       req.Price,
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
