package mocks

import (
	"github.com/getflight/flight/models"
	"github.com/stretchr/testify/mock"
)

type FileHelperMock struct {
	mock.Mock
}

func (m *FileHelperMock) Package(manifest models.Manifest) (string, error) {
	args := m.Called(manifest)

	return args.String(0), args.Error(1)
}

func (m *FileHelperMock) ReadFile(filename string) (string, error) {
	args := m.Called(filename)

	return args.String(0), args.Error(1)
}

func (m *FileHelperMock) WriteFile(value string, filename string) error {
	args := m.Called(value, filename)

	return args.Error(0)
}
