package mocks

import (
	"github.com/stretchr/testify/mock"
)

type TokenHelperMock struct {
	mock.Mock
}

func (m *TokenHelperMock) TokenExists() bool {
	args := m.Called()

	return args.Bool(0)
}

func (m *TokenHelperMock) SaveToken(token string) error {
	args := m.Called(token)

	return args.Error(0)
}

func (m *TokenHelperMock) GetToken() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}

func (m *TokenHelperMock) SaveOrganisation(organisationId string) error {
	args := m.Called(organisationId)

	return args.Error(0)
}

func (m *TokenHelperMock) GetOrganisation() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}
