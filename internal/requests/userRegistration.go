package requests

import (
	"errors"
	"net/mail"
)

type UserRegistration struct {
	Password string `json:"password"`
	Email    string `json:"email"`
	Name     string `json:"name"`
}

func (s *UserRegistration) Validate() error {
	if err := validateEmail(s.Email); err != nil {
		return err
	}

	if s.Password == "" {
		return errors.New("password cannot be empty")
	}

	if s.Name == "" {
		return errors.New("name cannot be empty")
	}

	return nil
}

func validateEmail(email string) error {
	_, err := mail.ParseAddress(email)

	return err
}
