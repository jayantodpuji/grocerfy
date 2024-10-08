// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MockUserHandler is an autogenerated mock type for the UserHandler type
type MockUserHandler struct {
	mock.Mock
}

type MockUserHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUserHandler) EXPECT() *MockUserHandler_Expecter {
	return &MockUserHandler_Expecter{mock: &_m.Mock}
}

// Register provides a mock function with given fields: _a0
func (_m *MockUserHandler) Register(_a0 echo.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Register")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUserHandler_Register_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Register'
type MockUserHandler_Register_Call struct {
	*mock.Call
}

// Register is a helper method to define mock.On call
//   - _a0 echo.Context
func (_e *MockUserHandler_Expecter) Register(_a0 interface{}) *MockUserHandler_Register_Call {
	return &MockUserHandler_Register_Call{Call: _e.mock.On("Register", _a0)}
}

func (_c *MockUserHandler_Register_Call) Run(run func(_a0 echo.Context)) *MockUserHandler_Register_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockUserHandler_Register_Call) Return(_a0 error) *MockUserHandler_Register_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUserHandler_Register_Call) RunAndReturn(run func(echo.Context) error) *MockUserHandler_Register_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUserHandler creates a new instance of MockUserHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUserHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUserHandler {
	mock := &MockUserHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
