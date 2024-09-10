package requests

import "net/mail"

type UserRegistration struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (s *UserRegistration) Validate() error {
	if err := validateEmail(s.Email); err != nil {
		return err
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	return err
}
