package service

import (
	"fmt"
	"github.com/getflight/flight/helpers"
	"github.com/getflight/flight/http"
	"github.com/getflight/flight/models"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/pkg/errors"
)

type LoginServiceType interface {
	Login(email string, password string) error
}

type LoginService struct {
	Client      http.ClientType
	TokenHelper helpers.TokenHelperType
}

func (s *LoginService) Login(email string, password string) error {
	login := models.Login{
		Email:    email,
		Password: password,
	}

	token, err := s.Client.Login(login)

	if err != nil {
		return errors.WithStack(err)
	}

	err = s.TokenHelper.SaveToken(token.Value)

	if err != nil {
		return errors.WithStack(err)
	}

	user, err := s.Client.GetUser()

	if err != nil {
		return errors.WithStack(err)
	}

	organisationId, err := s.getOrganisation(user)

	if err != nil {
		return errors.WithStack(err)
	}

	err = s.TokenHelper.SaveOrganisation(organisationId)

	if err != nil {
		return errors.WithStack(err)
	}

	log.Info("login successful")

	return nil
}

func (s *LoginService) getOrganisation(user models.User) (string, error) {
	if len(user.Organisations) == 0 {
		return "", errors.WithStack(errors.New("you need to be part of an organisation to login"))
	} else if len(user.Organisations) == 1 {
		return user.Organisations[0].ID, nil
	} else {
		fmt.Println("select organisation: ")
		for index, organisation := range user.Organisations {
			fmt.Println(fmt.Sprintf("%d. %s", index+1, organisation.Name))
		}

		return s.askForOrganisation(user)
	}
}

func (s *LoginService) askForOrganisation(user models.User) (string, error) {
	var input string

	_, err := fmt.Scanln(&input)

	if err != nil {
		return "", errors.WithStack(err)
	}

	index, err := strconv.Atoi(input)

	if err != nil || index <= 0 || index > len(user.Organisations) {
		fmt.Println("invalid choice, try again")
		return s.askForOrganisation(user)
	}

	return user.Organisations[index-1].ID, nil
}
