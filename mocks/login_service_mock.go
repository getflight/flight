package mocks

import "github.com/stretchr/testify/mock"

type LoginServiceMock struct {
	mock.Mock
}

func (m *LoginServiceMock) Login(email string, password string) error {
	args := m.Called(email, password)

	return args.Error(0)
}
