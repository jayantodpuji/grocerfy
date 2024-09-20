package services

import (
	"context"

	"github.com/gofrs/uuid"
	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/responses"
)

type GroceryListService interface {
	CreateGroceryList(context.Context, uuid.UUID, requests.CreateGroceryListRequest) error
	GetGroceryListByUserID(context.Context, uuid.UUID) ([]*responses.GroceryListIndexResponse, error)
	GetGroceryListByID(context.Context, uuid.UUID) (*responses.GroceryListDetailResponse, error)
}

type groceryListService struct {
	groceryListRepository     repositories.GroceryListRepository
	groceryListItemRepository repositories.GroceryListItemRepository
}

type GroceryListServiceDependency struct {
	GroceryListRepository     repositories.GroceryListRepository
	GroceryListItemRepository repositories.GroceryListItemRepository
}

func NewGroceryListService(deps GroceryListServiceDependency) GroceryListService {
	return &groceryListService{
		groceryListRepository:     deps.GroceryListRepository,
		groceryListItemRepository: deps.GroceryListItemRepository,
	}
}

func (g *groceryListService) CreateGroceryList(c context.Context, uid uuid.UUID, p requests.CreateGroceryListRequest) error {
	if err := g.groceryListRepository.InsertRecord(c, &models.GroceryList{
		Name:        p.Name,
		UserID:      uid,
		Description: p.Description,
	}); err != nil {
		return err
	}

	return nil
}

func (g *groceryListService) GetGroceryListByUserID(c context.Context, userID uuid.UUID) ([]*responses.GroceryListIndexResponse, error) {
	gcs, err := g.groceryListRepository.GetGroceryListByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	resps := make([]*responses.GroceryListIndexResponse, 0)
	for i := 0; i < len(gcs); i++ {
		resps = append(resps, &responses.GroceryListIndexResponse{
			ID:          gcs[i].ID.String(),
			Name:        gcs[i].Name,
			Description: gcs[i].Description,
			CreatedAt:   gcs[i].CreatedAt,
		})
	}

	return resps, nil
}

// TODO: optimize this
func (g *groceryListService) GetGroceryListByID(c context.Context, listID uuid.UUID) (*responses.GroceryListDetailResponse, error) {
	gl, err := g.groceryListRepository.GetGroceryListByID(c, listID)
	if err != nil {
		return nil, err
	}

	gli, err := g.groceryListItemRepository.GetItemsByGroceryList(c, gl.ID)
	if err != nil {
		return nil, err
	}

	items := make([]responses.GroceryListItemDetail, 0)
	for i := 0; i < len(gli); i++ {
		items = append(items, responses.GroceryListItemDetail{
			ID:            gli[i].ID,
			GroceryListID: gli[i].GroceryListID,
			Category:      gli[i].Category,
			Name:          gli[i].Name,
			Unit:          gli[i].Unit,
			Quantity:      gli[i].Quantity,
			Price:         gli[i].Price,
			IsPurchased:   gli[i].IsPurchased,
			CreatedAt:     gli[i].CreatedAt,
			UpdatedAt:     gli[i].UpdatedAt,
		})
	}

	detail := responses.GroceryListDetailResponse{
		ID:          gl.ID.String(),
		Name:        gl.Name,
		Description: gl.Description,
		CreatedAt:   gl.CreatedAt,
		Items:       items,
	}

	return &detail, nil
}
