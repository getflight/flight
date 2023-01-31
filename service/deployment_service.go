package service

import (
	"fmt"
	"github.com/getflight/flight/context"
	"github.com/getflight/flight/helpers"
	"github.com/getflight/flight/http"
	"github.com/getflight/flight/models"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/samber/lo"

	"github.com/pkg/errors"

	"time"

	log "github.com/sirupsen/logrus"
)

const (
	artifactPollRetries = 20
	stateCompleted      = "completed"
	stateExecuting      = "executing"
	stateFailed         = "failed"
	stateInitial        = "initial"
)

var (
	stepDescriptionMap = map[string]string{"ensure_artifact_exists": "ensuring artifact exists", "ensure_environment_configured": "ensuring environment is configured", "ensure_databases_exists": "ensuring databases exists", "create_database": "creating database", "deploy_artifact": "deploying artifact"}
)

type DeploymentServiceType interface {
	Deploy(environment string) error
}

type DeploymentService struct {
	Client        http.ClientType
	Configuration context.ConfigurationType
	FileHelper    helpers.FileHelperType
	TokenHelper   helpers.TokenHelperType
	start         time.Time
}

func (s *DeploymentService) Deploy(environment string) error {
	s.start = time.Now()

	err := s.verifyToken()

	if err != nil {
		return errors.WithStack(err)
	}

	log.Infof("deploying to %s", environment)

	err = s.initializeConfiguration()

	if err != nil {
		return errors.WithStack(err)
	}

	manifest, err := s.parseManifest()

	if err != nil {
		return errors.WithStack(err)
	}

	err = s.validateManifest(manifest, environment)

	if err != nil {
		return errors.WithStack(err)
	}

	content, err := s.packageArtifact(manifest)

	if err != nil {
		return errors.WithStack(err)
	}

	artifact, err := s.saveArtifact()

	if err != nil {
		return errors.WithStack(err)
	}

	artifact, err = s.pollArtifactForUpload(artifact)

	if err != nil {
		return errors.WithStack(err)
	}

	err = s.uploadArtifact(artifact, content)

	if err != nil {
		return errors.WithStack(err)
	}

	deployment, err := s.saveDeployment(artifact, environment, manifest)

	if err != nil {
		return errors.WithStack(err)
	}

	err = s.pollDeployment(deployment)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *DeploymentService) verifyToken() error {
	if !s.TokenHelper.TokenExists() {
		return errors.New("token not found, login to deploy")
	}

	return nil
}

func (s *DeploymentService) initializeConfiguration() error {
	if s.Configuration != nil {
		return nil
	}

	config := &context.Configuration{}
	err := config.Init()

	if err != nil {
		return errors.WithStack(err)
	}

	s.Configuration = config

	return nil
}

func (s *DeploymentService) parseManifest() (models.Manifest, error) {
	log.Info("parsing manifest")
	manifest, err := s.Configuration.GetManifest()

	if err != nil {
		return manifest, errors.WithStack(err)
	}

	return manifest, nil
}

func (s *DeploymentService) validateManifest(manifest models.Manifest, environment string) error {
	err := validator.New().Struct(manifest)

	if err != nil {
		return errors.WithStack(err)
	}

	if !lo.ContainsBy[models.ManifestEnvironment](manifest.Environments, func(manifestEnvironment models.ManifestEnvironment) bool {
		return manifestEnvironment.Name == environment
	}) {
		return errors.New(fmt.Sprintf("environment %s not found in flight.yml, please configure environment before deploying", environment))
	}

	return nil
}

func (s *DeploymentService) packageArtifact(manifest models.Manifest) (string, error) {
	log.Info("packaging artifact")
	content, err := s.FileHelper.Package(manifest)

	if err != nil {
		return content, errors.WithStack(err)
	}

	return content, nil
}

func (s *DeploymentService) saveArtifact() (models.Artifact, error) {
	log.Info("saving artifact")
	artifact := models.Artifact{}
	artifact, err := s.Client.SaveArtifact(artifact)

	if err != nil {
		return artifact, errors.WithStack(err)
	}

	return artifact, nil
}

func (s *DeploymentService) pollArtifactForUpload(artifact models.Artifact) (models.Artifact, error) {
	log.Info("polling artifact for upload")
	time.Sleep(time.Second)

	for i := 0; i < artifactPollRetries; i++ {
		artifact, err := s.Client.GetArtifact(artifact.ID)

		if err != nil {
			return artifact, errors.WithStack(err)
		}

		if artifact.UploadURL != "" {
			return artifact, nil
		}

		time.Sleep(time.Second)
		log.Debugf("artifact preparing for upload pending, retry attempt: %v for artifact: %v", i+1, artifact.ID)
	}

	return artifact, errors.WithStack(errors.New(fmt.Sprintf("artifact failed to prepare for upload: %s", artifact.ID)))
}

func (s *DeploymentService) uploadArtifact(artifact models.Artifact, content string) error {
	log.Info("uploading artifact")
	err := s.Client.UploadArtifact(artifact, content)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *DeploymentService) saveDeployment(artifact models.Artifact, environment string, manifest models.Manifest) (models.Deployment, error) {
	log.Info("initiating deployment")
	deployment := models.Deployment{
		Artifact:    artifact.ID,
		Environment: environment,
		Manifest:    manifest,
	}
	deployment, err := s.Client.SaveDeployment(deployment)

	if err != nil {
		return deployment, errors.WithStack(err)
	}

	return deployment, nil
}

func (s *DeploymentService) pollDeployment(deployment models.Deployment) error {
	deployment, err := s.printDeploymentSteps(deployment)

	if err != nil {
		return errors.WithStack(err)
	}

	if deployment.State == stateCompleted {
		log.Infof("deployment #%s completed successfully in %s", deployment.Count, time.Since(s.start).Round(time.Second))
	} else {
		step, found := lo.Find[models.DeploymentStep](deployment.Steps, func(step models.DeploymentStep) bool {
			return strings.EqualFold(step.State, stateFailed)
		})

		if found {
			log.Infof("deployment #%s failed on step %s with error %s", deployment.Count, step.Name, step.Result)
		} else {
			log.Infof("deployment #%s failed", deployment.Count)
		}
	}

	return nil
}

func (s *DeploymentService) printDeploymentSteps(deployment models.Deployment) (models.Deployment, error) {
	var printedSteps []string
	done := false

	for !done {

		var err error
		deployment, err = s.Client.GetDeployment(deployment.ID)

		if err != nil {
			return deployment, errors.WithStack(err)
		}

		for _, step := range deployment.Steps {
			if !lo.Contains(printedSteps, step.ID) && (step.State == stateCompleted || step.State == stateFailed || step.State == stateExecuting) {
				printedSteps = append(printedSteps, step.ID)

				stepDescription := stepDescriptionMap[step.Name]

				if stepDescription == "" {
					stepDescription = step.Name
				}

				log.Info(stepDescription)
			}
		}

		done = deployment.State != stateInitial && deployment.State != stateExecuting

		time.Sleep(time.Second)
	}

	return deployment, nil
}

func (s *DeploymentService) getProject(manifestEnvironment string, manifestProject string) (*models.Project, error) {
	organisationId, err := s.TokenHelper.GetOrganisation()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	organisation, err := s.Client.GetOrganisation(organisationId)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	environment, found := lo.Find[models.Environment](organisation.Environments, func(environment models.Environment) bool {
		return strings.EqualFold(environment.Name, manifestEnvironment)
	})

	if !found {
		return nil, errors.WithStack(errors.New(fmt.Sprintf("environment not found in organisation: %s", manifestEnvironment)))
	}

	environment, err = s.Client.GetEnvironment(environment.ID)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	project, found := lo.Find[models.Project](environment.Projects, func(project models.Project) bool {
		return strings.EqualFold(project.Name, manifestProject)
	})

	if !found {
		return nil, errors.WithStack(errors.New(fmt.Sprintf("project not found in organisation: %s", manifestProject)))
	}

	return &project, nil
}
