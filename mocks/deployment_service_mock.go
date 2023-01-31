package mocks

import "github.com/stretchr/testify/mock"

type DeploymentServiceMock struct {
	mock.Mock
}

func (m *DeploymentServiceMock) Deploy(environment string) error {
	args := m.Called(environment)

	return args.Error(0)
}
