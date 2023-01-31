package helpers

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type TokenHelperType interface {
	TokenExists() bool
	SaveToken(token string) error
	GetToken() (string, error)
	SaveOrganisation(organisationId string) error
	GetOrganisation() (string, error)
}

type TokenHelper struct {
	FileHelper FileHelperType
}

func (h *TokenHelper) TokenExists() bool {
	token, err := h.GetToken()

	if err != nil {
		log.Debugf("%+v", err)

		return false
	}

	return token != ""
}

func (h *TokenHelper) SaveToken(token string) error {
	err := h.FileHelper.WriteFile(token, tokenFilename)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (h *TokenHelper) GetToken() (string, error) {
	token, err := h.FileHelper.ReadFile(tokenFilename)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return token, nil
}

func (h *TokenHelper) SaveOrganisation(organisationId string) error {
	err := h.FileHelper.WriteFile(organisationId, organisationFilename)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (h *TokenHelper) GetOrganisation() (string, error) {
	organisationId, err := h.FileHelper.ReadFile(organisationFilename)

	if err != nil {
		log.Debugf("%+v", err)

		return "", nil
	}

	return organisationId, nil
}
