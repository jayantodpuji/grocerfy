package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jayantodpuji/grocerfy/internal/handlers"
	mocks "github.com/jayantodpuji/grocerfy/test/mocks/services"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	userService *mocks.MockUserService
	userHandler handlers.UserHandler
}

func (s *UserHandlerTestSuite) SetupTest() {
	s.userService = mocks.NewMockUserService(s.T())
	s.userHandler = handlers.NewUserHandler(handlers.UserHandlerDependency{UserService: s.userService})
}

func (s *UserHandlerTestSuite) Test_Register() {
	e := echo.New()
	p := `{"email": "john.doe@example.com", "password": "password123", "name": "John Doe"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(p))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	s.userService.EXPECT().Register(c.Request().Context(), mock.Anything).Return(nil)
	err := s.userHandler.Register(c)
	s.NoError(err)
	s.Equal(http.StatusOK, rec.Code)
}

func TestUserHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
