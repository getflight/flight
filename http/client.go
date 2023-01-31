package http

import (
	"fmt"
	"github.com/getflight/flight/helpers"
	"github.com/getflight/flight/models"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"github.com/imroc/req"
)

var (
	CustomApiUrl = ""
)

const (
	defaultApiUrl = "https://api.getflight.io"
	apiVersion    = "/v1"
)

type ClientType interface {
	GetArtifact(artifactID string) (models.Artifact, error)
	SaveArtifact(artifact models.Artifact) (models.Artifact, error)
	UploadArtifact(artifact models.Artifact, content string) error
	SaveDeployment(deployment models.Deployment) (models.Deployment, error)
	GetDeployment(deploymentID string) (models.Deployment, error)
	Login(login models.Login) (models.Token, error)
	GetUser() (models.User, error)
	GetOrganisation(organisationId string) (models.Organisation, error)
	GetEnvironment(environmentId string) (models.Environment, error)
	GetProject(projectId string) (models.Project, error)
}

type Client struct {
	TokenHelper *helpers.TokenHelper
}

func (c *Client) GetArtifact(artifactID string) (models.Artifact, error) {
	artifact := &models.Artifact{}
	headers, err := c.getHeaders()

	if err != nil {
		return *artifact, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/artifacts/%s", artifactID), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *artifact, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *artifact, errors.WithStack(err)
	}

	err = r.ToJSON(artifact)

	if err != nil {
		return *artifact, errors.WithStack(err)
	}

	return *artifact, nil
}

func (c *Client) SaveArtifact(artifact models.Artifact) (models.Artifact, error) {
	headers, err := c.getHeaders()

	if err != nil {
		return artifact, errors.WithStack(err)
	}

	r, err := req.Post(c.getUrl("/artifacts"), headers, req.BodyJSON(&artifact))

	log.Debugf("%+v", r)

	if err != nil {
		return artifact, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return artifact, errors.WithStack(err)
	}

	err = r.ToJSON(&artifact)

	if err != nil {
		return artifact, errors.WithStack(err)
	}

	return artifact, nil
}

// UploadArtifact uploads to the storage provider. The upload URL is
// provided by the api and returned after the SaveArtifact call.
func (c *Client) UploadArtifact(artifact models.Artifact, content string) error {
	r, err := req.Put(artifact.UploadURL, content)

	if err != nil {
		return errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return errors.WithStack(err)
	}

	return nil
}

// SaveDeployment provisions the artifact and deploys it on the serverless infrastructure.
// The artifact must be uploaded on the cloud storage before deploying
func (c *Client) SaveDeployment(deployment models.Deployment) (models.Deployment, error) {

	headers, err := c.getHeaders()

	if err != nil {
		return deployment, errors.WithStack(err)
	}

	req.SetTimeout(5 * time.Minute)
	r, err := req.Post(c.getUrl("/deployments"), headers, req.BodyJSON(&deployment))
	req.SetTimeout(30 * time.Second)

	log.Debugf("%+v", r)

	if err != nil {
		return deployment, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return deployment, errors.WithStack(err)
	}

	err = r.ToJSON(&deployment)

	if err != nil {
		return deployment, errors.WithStack(err)
	}

	return deployment, nil
}

func (c *Client) GetDeployment(deploymentID string) (models.Deployment, error) {
	deployment := &models.Deployment{}
	headers, err := c.getHeaders()

	if err != nil {
		return *deployment, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/deployments/%s", deploymentID), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *deployment, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *deployment, errors.WithStack(err)
	}

	err = r.ToJSON(deployment)

	if err != nil {
		return *deployment, errors.WithStack(err)
	}

	return *deployment, nil
}

func (c *Client) Login(login models.Login) (models.Token, error) {
	token := &models.Token{}
	headers := req.Header{
		"Accept": "application/json",
	}

	r, err := req.Post(c.getUrl("/auth/login"), headers, req.BodyJSON(&login))

	log.Debugf("%+v", r)

	if err != nil {
		return *token, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *token, errors.WithStack(err)
	}

	err = r.ToJSON(token)

	if err != nil {
		return *token, errors.WithStack(err)
	}

	return *token, nil
}

func (c *Client) GetUser() (models.User, error) {
	user := &models.User{}
	headers, err := c.getHeaders()

	if err != nil {
		return *user, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/me"), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *user, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *user, errors.WithStack(err)
	}

	err = r.ToJSON(user)

	if err != nil {
		return *user, errors.WithStack(err)
	}

	return *user, nil
}

func (c *Client) GetOrganisation(organisationId string) (models.Organisation, error) {
	organisation := &models.Organisation{}
	headers, err := c.getHeaders()

	if err != nil {
		return *organisation, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/organisations/%s", organisationId), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *organisation, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *organisation, errors.WithStack(err)
	}

	err = r.ToJSON(organisation)

	if err != nil {
		return *organisation, errors.WithStack(err)
	}

	return *organisation, nil
}

func (c *Client) GetEnvironment(environmentId string) (models.Environment, error) {
	environment := &models.Environment{}
	headers, err := c.getHeaders()

	if err != nil {
		return *environment, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/environments/%s", environmentId), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *environment, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *environment, errors.WithStack(err)
	}

	err = r.ToJSON(environment)

	if err != nil {
		return *environment, errors.WithStack(err)
	}

	return *environment, nil
}

func (c *Client) GetProject(projectId string) (models.Project, error) {
	project := &models.Project{}
	headers, err := c.getHeaders()

	if err != nil {
		return *project, errors.WithStack(err)
	}

	r, err := req.Get(c.getUrl("/projects/%s", projectId), headers)

	log.Debugf("%+v", r)

	if err != nil {
		return *project, errors.WithStack(err)
	}

	if !c.isOk(r) {
		err = c.handleError(r)

		return *project, errors.WithStack(err)
	}

	err = r.ToJSON(project)

	if err != nil {
		return *project, errors.WithStack(err)
	}

	return *project, nil
}

func (c *Client) getUrl(path string, args ...any) string {
	url := defaultApiUrl

	if CustomApiUrl != "" {
		url = CustomApiUrl
	}

	url += apiVersion

	if len(args) > 0 {
		url += fmt.Sprintf(path, args...)
	} else {
		url += path
	}

	return url
}

func (c *Client) getHeaders() (req.Header, error) {
	var header req.Header
	token, err := c.TokenHelper.GetToken()

	if err != nil {
		return header, errors.WithStack(err)
	}

	header = req.Header{
		"Accept":        "application/json",
		"Authorization": "Bearer " + token,
	}

	organisation, err := c.TokenHelper.GetOrganisation()

	if err != nil {
		return header, errors.WithStack(err)
	}

	if organisation != "" {
		header["x-flight-organisation-id"] = organisation
	}

	return header, nil
}

func (c *Client) isOk(resp *req.Resp) bool {
	return resp.Response().StatusCode >= 200 && resp.Response().StatusCode <= 299
}

func (c *Client) handleError(r *req.Resp) error {
	log.Debugf("API error %d %s %s", r.Response().StatusCode, r.Request().Method, r.Request().URL)

	errorResponse := &models.Error{}
	err := r.ToJSON(errorResponse)

	if err != nil {
		return errors.New(fmt.Sprintf("API error http code %d", r.Response().StatusCode))
	}

	return errors.New(errorResponse.Message)
}
