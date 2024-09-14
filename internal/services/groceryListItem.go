package services

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
)

type GroceryListItemService interface {
	CreateGroceryListItem(c context.Context, req *requests.CreateGroceryListItemRequest) (*models.GroceryListItem, error)
	BulkCreateGroceryListItem(c context.Context, req []*requests.CreateGroceryListItemRequest) ([]*models.GroceryListItem, error)
}

type groceryListItemService struct {
	repo repositories.GroceryListItemRepository
}

func NewGroceryListItemService(repo repositories.GroceryListItemRepository) GroceryListItemService {
	return &groceryListItemService{repo: repo}
}

func (s *groceryListItemService) CreateGroceryListItem(c context.Context, req *requests.CreateGroceryListItemRequest) (*models.GroceryListItem, error) {
	gli, err := s.repo.InsertRecord(c, &models.GroceryListItem{
		GroceryListID: uuid.Must(uuid.FromString(req.GroceryListID)),
		Category:      req.Category,
		Name:          req.Name,
		Unit:          req.Unit,
		Size:          req.Size,
		Quantity:      req.Quantity,
		Price:         req.Price,
	})

	if err != nil {
		return nil, err
	}

	return gli, nil
}

func (s *groceryListItemService) BulkCreateGroceryListItem(c context.Context, req []*requests.CreateGroceryListItemRequest) ([]*models.GroceryListItem, error) {
	gli := make([]*models.GroceryListItem, 0)

	for _, r := range req {
		gli = append(gli, &models.GroceryListItem{
			GroceryListID: uuid.Must(uuid.FromString(r.GroceryListID)),
			Category:      r.Category,
			Name:          r.Name,
			Unit:          r.Unit,
			Size:          r.Size,
			Quantity:      r.Quantity,
			Price:         r.Price,
		})
	}

	return s.repo.InsertRecords(c, gli)
}
