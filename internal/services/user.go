package services

import (
	"context"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/internal/requests"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(context.Context, requests.UserRegistration) error
	Login(context.Context, requests.UserLogin) (string, error)
}

type userService struct {
	userRepository repositories.UserRepository
	authService    AuthService
}

type UserServiceDependency struct {
	UserRepository repositories.UserRepository
	AuthService    AuthService
}

func NewUserService(deps UserServiceDependency) UserService {
	return &userService{userRepository: deps.UserRepository, authService: deps.AuthService}
}

func (us *userService) Register(c context.Context, p requests.UserRegistration) error {
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

func (us *userService) Login(c context.Context, p requests.UserLogin) (string, error) {
	user, err := us.userRepository.FindUserByEmail(c, p.Email)
	if err != nil {
		return "", err
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(p.Password)); err != nil {
		return "", err
	}

	token, err := us.authService.GenerateToken(c, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}
