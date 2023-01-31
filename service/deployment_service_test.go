package service

import (
	"github.com/getflight/flight/mocks"
	"github.com/getflight/flight/models"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeploymentService(t *testing.T) {
	t.Run("verifyToken with token returns nil", func(t *testing.T) {
		// given
		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("TokenExists").Return(true)

		deploymentService := DeploymentService{
			TokenHelper: tokenHelperMock,
		}

		// when
		result := deploymentService.verifyToken()

		// then
		assert.Nil(t, result)
		tokenHelperMock.AssertExpectations(t)
	})

	t.Run("verifyToken without token returns error", func(t *testing.T) {
		// given
		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("TokenExists").Return(false)

		deploymentService := DeploymentService{
			TokenHelper: tokenHelperMock,
		}

		// when
		result := deploymentService.verifyToken()

		// then
		assert.NotNil(t, result)
		tokenHelperMock.AssertExpectations(t)
	})

	t.Run("parseManifest with manifest returns manifest", func(t *testing.T) {
		// given
		configuration := &mocks.ConfigurationMock{}
		configuration.On("GetManifest").Return(models.Manifest{}, nil)

		deploymentService := DeploymentService{
			Configuration: configuration,
		}

		// when
		manifest, err := deploymentService.parseManifest()

		// then
		assert.NotNil(t, manifest)
		assert.Nil(t, err)
	})

	t.Run("parseManifest with error returns error", func(t *testing.T) {
		// given
		configuration := &mocks.ConfigurationMock{}
		configuration.On("GetManifest").Return(models.Manifest{}, errors.New("test error"))

		deploymentService := DeploymentService{
			Configuration: configuration,
		}

		// when
		_, err := deploymentService.parseManifest()

		// then
		assert.NotNil(t, err)
	})

	t.Run("validateManifest with valid manifest returns nil", func(t *testing.T) {
		// given
		manifest := getManifest()

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.Nil(t, err)
	})

	t.Run("validateManifest with invalid environment returns error", func(t *testing.T) {
		// given
		manifest := getManifest()

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "test")

		// then
		assert.NotNil(t, err)
	})

	t.Run("validateManifest with invalid name returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Name = ""

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Name", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid trigger returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Trigger = "test"

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Trigger", validationErrors[0].Field())
	})

	t.Run("validateManifest with empty environments returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments = []models.ManifestEnvironment{}

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Environments", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid environment name returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments[0].Name = ""

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Name", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid database name returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments[0].Databases[0].Name = ""

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Name", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid database driver returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments[0].Databases[0].Driver = "test"

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Driver", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid variable key returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments[0].Variables[0].Key = ""

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Key", validationErrors[0].Field())
	})

	t.Run("validateManifest with invalid variable value returns error", func(t *testing.T) {
		// given
		manifest := getManifest()
		manifest.Environments[0].Variables[0].Value = ""

		deploymentService := DeploymentService{}

		// when
		err := deploymentService.validateManifest(manifest, "dev")

		// then
		assert.NotNil(t, err)
		validationErrors := errors.Unwrap(err).(validator.ValidationErrors)
		assert.Equal(t, 1, len(validationErrors))
		assert.Equal(t, "Value", validationErrors[0].Field())
	})

	t.Run("packageArtifact with success returns file content", func(t *testing.T) {
		// given
		manifest := getManifest()

		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("Package", manifest).Return("content", nil)

		deploymentService := DeploymentService{
			FileHelper: fileHelperMock,
		}

		// when
		content, err := deploymentService.packageArtifact(manifest)

		// then
		assert.Nil(t, err)
		assert.Equal(t, "content", content)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("packageArtifact with error returns error", func(t *testing.T) {
		// given
		manifest := getManifest()

		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("Package", manifest).Return("", errors.New("test error"))

		deploymentService := DeploymentService{
			FileHelper: fileHelperMock,
		}

		// when
		content, err := deploymentService.packageArtifact(manifest)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "", content)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("saveArtifact with success returns artifact and nil", func(t *testing.T) {
		// given
		clientMock := &mocks.ClientMock{}
		clientMock.On("SaveArtifact", mock.Anything).Return(models.Artifact{}, nil)

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		artifact, err := deploymentService.saveArtifact()

		// then
		assert.Nil(t, err)
		assert.NotNil(t, artifact)
		clientMock.AssertExpectations(t)
	})

	t.Run("saveArtifact with error returns error", func(t *testing.T) {
		// given
		clientMock := &mocks.ClientMock{}
		clientMock.On("SaveArtifact", mock.Anything).Return(models.Artifact{}, errors.New("test error"))

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		_, err := deploymentService.saveArtifact()

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("pollArtifactForUpload polls until uploadUrl is set", func(t *testing.T) {
		// given
		artifact := models.Artifact{
			ID: "1",
		}

		artifactWithUploadUrl := models.Artifact{
			ID:        "1",
			UploadURL: "test",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetArtifact", "1").Return(artifact, nil).Once()
		clientMock.On("GetArtifact", "1").Return(artifact, nil).Once()
		clientMock.On("GetArtifact", "1").Return(artifactWithUploadUrl, nil).Once()

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		artifact, err := deploymentService.pollArtifactForUpload(artifact)

		// then
		assert.Nil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("pollArtifactForUpload with error returns error", func(t *testing.T) {
		// given
		artifact := models.Artifact{
			ID: "1",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetArtifact", "1").Return(artifact, errors.New("test error")).Once()

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		artifact, err := deploymentService.pollArtifactForUpload(artifact)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("uploadArtifact with success returns nil", func(t *testing.T) {
		// given
		artifact := models.Artifact{}
		content := "content"

		clientMock := &mocks.ClientMock{}
		clientMock.On("UploadArtifact", artifact, content).Return(nil)

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		err := deploymentService.uploadArtifact(artifact, content)

		// then
		assert.Nil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("uploadArtifact with error returns error", func(t *testing.T) {
		// given
		artifact := models.Artifact{}
		content := "content"

		clientMock := &mocks.ClientMock{}
		clientMock.On("UploadArtifact", artifact, content).Return(errors.New("test error"))

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		err := deploymentService.uploadArtifact(artifact, content)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("saveDeployment with success returns deployment and nil", func(t *testing.T) {
		// given
		artifact := models.Artifact{}
		environment := "dev"
		manifest := getManifest()

		clientMock := &mocks.ClientMock{}
		clientMock.On("SaveDeployment", mock.Anything).Return(models.Deployment{}, nil)

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		deployment, err := deploymentService.saveDeployment(artifact, environment, manifest)

		// then
		assert.Nil(t, err)
		assert.NotNil(t, deployment)
		clientMock.AssertExpectations(t)
	})

	t.Run("saveDeployment with error returns error", func(t *testing.T) {
		// given
		artifact := models.Artifact{}
		environment := "dev"
		manifest := getManifest()

		clientMock := &mocks.ClientMock{}
		clientMock.On("SaveDeployment", mock.Anything).Return(models.Deployment{}, errors.New("test error"))

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		_, err := deploymentService.saveDeployment(artifact, environment, manifest)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("printDeploymentSteps polls deployment until completed", func(t *testing.T) {
		// given
		deployment := models.Deployment{
			ID:       "1",
			State:    stateInitial,
			Manifest: models.Manifest{},
			Steps: []models.DeploymentStep{
				{
					ID:    "1",
					Name:  "test",
					State: stateCompleted,
				},
			},
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetDeployment", "1").Return(deployment, nil).Once()

		deployment.State = stateCompleted

		clientMock.On("GetDeployment", "1").Return(deployment, nil).Once()

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		result, err := deploymentService.printDeploymentSteps(deployment)

		// then
		assert.Nil(t, err)
		assert.Equal(t, deployment, result)
		clientMock.AssertExpectations(t)
	})

	t.Run("printDeploymentSteps with error return error", func(t *testing.T) {
		// given
		deployment := models.Deployment{
			ID: "1",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetDeployment", "1").Return(deployment, errors.New("test error"))

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		_, err := deploymentService.printDeploymentSteps(deployment)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("pollDeployment with error return error", func(t *testing.T) {
		// given
		deployment := models.Deployment{
			ID: "1",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetDeployment", "1").Return(deployment, errors.New("test error"))

		deploymentService := DeploymentService{
			Client: clientMock,
		}

		// when
		err := deploymentService.pollDeployment(deployment)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("getProject finds project from environment", func(t *testing.T) {
		// given
		manifestEnvironment := "dev"
		manifestProject := "app"
		organisation := models.Organisation{
			ID:   "1",
			Name: "org",
			Environments: []models.Environment{
				{
					ID:   "2",
					Name: "dev",
				},
			},
		}

		environment := models.Environment{
			ID:   "2",
			Name: "dev",
			Projects: []models.Project{
				{
					ID:   "3",
					Name: "app",
				},
			},
		}

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("GetOrganisation").Return("1", nil)

		clientMock := &mocks.ClientMock{}
		clientMock.On("GetOrganisation", "1").Return(organisation, nil)
		clientMock.On("GetEnvironment", "2").Return(environment, nil)

		deploymentService := DeploymentService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		project, err := deploymentService.getProject(manifestEnvironment, manifestProject)

		// then
		assert.Nil(t, err)
		assert.NotNil(t, project)
		assert.Equal(t, "3", project.ID)
		clientMock.AssertExpectations(t)
	})

	t.Run("Deploy calls expected and returns nil", func(t *testing.T) {
		// given
		environment := "dev"
		artifact := models.Artifact{
			ID:        "1",
			UploadURL: "test",
		}
		content := "content"
		deployment := models.Deployment{
			ID:       "1",
			State:    stateCompleted,
			Manifest: models.Manifest{},
			Steps: []models.DeploymentStep{
				{
					ID:    "1",
					Name:  "test",
					State: stateCompleted,
				},
			},
		}
		manifest := getManifest()

		clientMock := &mocks.ClientMock{}
		clientMock.On("SaveArtifact", mock.Anything).Return(artifact, nil).Once()
		clientMock.On("GetArtifact", "1").Return(artifact, nil).Once()
		clientMock.On("UploadArtifact", artifact, content).Return(nil)
		clientMock.On("SaveDeployment", mock.Anything).Return(deployment, nil).Once()
		clientMock.On("GetDeployment", "1").Return(deployment, nil).Once()

		configuration := &mocks.ConfigurationMock{}
		configuration.On("GetManifest").Return(manifest, nil)

		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("Package", manifest).Return("content", nil)

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("TokenExists").Return(true)

		deploymentService := DeploymentService{
			Client:        clientMock,
			Configuration: configuration,
			FileHelper:    fileHelperMock,
			TokenHelper:   tokenHelperMock,
		}

		// when
		err := deploymentService.Deploy(environment)

		// then
		assert.Nil(t, err)
		clientMock.AssertExpectations(t)
	})
}

func getManifest() models.Manifest {
	files := &[]string{
		"file1",
		"file2",
		"file3",
	}

	manifest := models.Manifest{
		Name:    "test",
		Files:   files,
		Trigger: "queue",
		Environments: []models.ManifestEnvironment{
			{
				Name: "dev",
				Databases: []models.ManifestDatabase{
					{
						Name:   "db",
						Driver: "mysql",
					},
				},
				Variables: []models.ManifestVariable{
					{
						Key:   "var1",
						Value: "value1",
					},
					{
						Key:   "var2",
						Value: "value2",
					},
				},
			},
		},
	}

	return manifest
}
