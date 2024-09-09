package services

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	SignUp(context.Context, requests.SignUp) error
}

type userService struct {
	userRepository repositories.UserRepository
}

type UserServiceDependency struct {
	UserRepository repositories.UserRepository
}

func NewUserService(deps UserServiceDependency) UserService {
	return &userService{userRepository: deps.UserRepository}
}

func (us *userService) SignUp(c context.Context, p requests.SignUp) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), 8)
	if err != nil {
		return err
	}

	nu := models.User{
		Email:        p.Email,
		PasswordHash: string(hashedPassword),
		Name:         p.Name,
	}

	if err = us.userRepository.InsertRecord(c, &nu); err != nil {
		return err
	}

	return nil
}
