package repositories_test

import (
	"context"
	"log"
	"testing"

	"github.com/jayantodpuji/grocerfy/internal/models"
	"github.com/jayantodpuji/grocerfy/internal/repositories"
	"github.com/jayantodpuji/grocerfy/test"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	db   *gorm.DB
	repo repositories.UserRepository
}

func (s *UserRepositoryTestSuite) SetupSuite() {
	db := test.TestDBInstance
	s.db = db
}

func (s *UserRepositoryTestSuite) SetupTest() {
	err := s.db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error
	if err != nil {
		log.Fatal(err)
	}

	s.repo = repositories.NewUserRepository(repositories.UserRepositoryDependency{DB: s.db})
}

func (s *UserRepositoryTestSuite) Test_InsertRecord() {
	nu := &models.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password123",
	}

	err := s.repo.InsertRecord(context.Background(), nu)
	s.NoError(err)
	s.NotZero(nu.ID)

	var u models.User
	err = s.db.First(&u, nu.ID).Error
	s.NoError(err)
	s.Equal(nu.ID, u.ID)
	s.Equal(nu.Name, u.Name)
	s.Equal(nu.Email, u.Email)
	s.Equal(nu.PasswordHash, u.PasswordHash)
}

func (s *UserRepositoryTestSuite) Test_FindUserByEmail() {
	nu := &models.User{
		Name:         "John Doe",
		Email:        "john.doe@example.com",
		PasswordHash: "password123",
	}

	err := s.repo.InsertRecord(context.Background(), nu)
	s.NoError(err)
	s.NotZero(nu.ID)

	eu, err := s.repo.FindUserByEmail(context.Background(), nu.Email)
	s.NoError(err)
	s.NotNil(eu)
	s.Equal(nu.ID, eu.ID)
	s.Equal(nu.Name, eu.Name)
	s.Equal(nu.Email, eu.Email)
	s.Equal(nu.PasswordHash, eu.PasswordHash)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
