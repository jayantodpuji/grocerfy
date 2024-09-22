package requests

type UpdateGroceryListRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
