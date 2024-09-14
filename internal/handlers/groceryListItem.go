package handlers

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
)

type GroceryListItemHandler interface {
	Create(c context.Context, req *requests.CreateGroceryListItemRequest) error
	BulkCreate(c context.Context, req []*requests.CreateGroceryListItemRequest) error
}

type groceryListItemHandler struct {
	service services.GroceryListItemService
}

type GroceryListItemHandlerDependency struct {
	Service services.GroceryListItemService
}

func NewGroceryListItemHandler(deps GroceryListItemHandlerDependency) GroceryListItemHandler {
	return &groceryListItemHandler{service: deps.Service}
}

func (h *groceryListItemHandler) Create(c context.Context, req *requests.CreateGroceryListItemRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	_, err := h.service.CreateGroceryListItem(c, req)
	if err != nil {
		return err
	}

	return nil
}

func (h *groceryListItemHandler) BulkCreate(c context.Context, req []*requests.CreateGroceryListItemRequest) error {
	_, err := h.service.BulkCreateGroceryListItem(c, req)
	if err != nil {
		return err
	}

	return nil
}
