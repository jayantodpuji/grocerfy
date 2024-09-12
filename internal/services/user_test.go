package services_test

import (
	"context"
	"testing"

	"github.com/jayantodpuji/grocerfy/internal/requests"
	"github.com/jayantodpuji/grocerfy/internal/services"
	mocks "github.com/jayantodpuji/grocerfy/test/mocks/repositories"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	userRepository *mocks.MockUserRepository
	userService    services.UserService
}

func (s *UserServiceTestSuite) SetupTest() {
	s.userRepository = mocks.NewMockUserRepository(s.T())
	s.userService = services.NewUserService(services.UserServiceDependency{UserRepository: s.userRepository})
}

func (s *UserServiceTestSuite) Test_Register() {
	ctx := context.Background()
	req := requests.UserRegistration{
		Email:    "john.doe@example.com",
		Password: "password123",
		Name:     "John Doe",
	}

	s.userRepository.EXPECT().InsertRecord(ctx, mock.Anything).Return(nil)

	err := s.userService.Register(ctx, req)
	s.NoError(err)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
