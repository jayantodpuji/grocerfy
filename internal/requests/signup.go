package requests

type SignUp struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}
