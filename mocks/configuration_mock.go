package mocks

import (
	"github.com/getflight/flight/models"
	"github.com/stretchr/testify/mock"
)

type ConfigurationMock struct {
	mock.Mock
}

func (m *ConfigurationMock) Init() error {
	args := m.Called()

	return args.Error(0)
}

func (m *ConfigurationMock) GetManifest() (models.Manifest, error) {
	args := m.Called()

	return args.Get(0).(models.Manifest), args.Error(1)
}
