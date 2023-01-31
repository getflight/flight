package mocks

import (
	"github.com/getflight/flight/models"
	"github.com/stretchr/testify/mock"
)

type ClientMock struct {
	mock.Mock
}

func (m *ClientMock) GetArtifact(artifactID string) (models.Artifact, error) {
	args := m.Called(artifactID)

	return args.Get(0).(models.Artifact), args.Error(1)
}

func (m *ClientMock) SaveArtifact(artifact models.Artifact) (models.Artifact, error) {
	args := m.Called(artifact)

	return args.Get(0).(models.Artifact), args.Error(1)
}

func (m *ClientMock) UploadArtifact(artifact models.Artifact, content string) error {
	args := m.Called(artifact, content)

	return args.Error(0)
}

func (m *ClientMock) SaveDeployment(deployment models.Deployment) (models.Deployment, error) {
	args := m.Called(deployment)

	return args.Get(0).(models.Deployment), args.Error(1)
}

func (m *ClientMock) GetDeployment(deploymentID string) (models.Deployment, error) {
	args := m.Called(deploymentID)

	return args.Get(0).(models.Deployment), args.Error(1)
}

func (m *ClientMock) Login(login models.Login) (models.Token, error) {
	args := m.Called(login)

	return args.Get(0).(models.Token), args.Error(1)
}

func (m *ClientMock) GetUser() (models.User, error) {
	args := m.Called()

	return args.Get(0).(models.User), args.Error(1)
}

func (m *ClientMock) GetOrganisation(organisationId string) (models.Organisation, error) {
	args := m.Called(organisationId)

	return args.Get(0).(models.Organisation), args.Error(1)
}

func (m *ClientMock) GetEnvironment(environmentId string) (models.Environment, error) {
	args := m.Called(environmentId)

	return args.Get(0).(models.Environment), args.Error(1)
}

func (m *ClientMock) GetProject(projectId string) (models.Project, error) {
	args := m.Called(projectId)

	return args.Get(0).(models.Project), args.Error(1)
}
