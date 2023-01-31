package commands

import (
	"fmt"
	"github.com/getflight/flight/service"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type Login struct {
	LoginService service.LoginServiceType
}

func (l *Login) command() *cobra.Command {
	var email string
	var password string

	command := &cobra.Command{
		Use:   "login",
		Short: "Login to authenticate with the api",
		Long:  `Authentication is necessary in order to run most commands`,
		Run: func(cmd *cobra.Command, args []string) {

			if password == "" {
				pwd, err := l.askForPassword()

				if err != nil {
					log.Debugf("%+v", err)
					log.Fatal(err.Error())
				}

				password = pwd
			}

			err := l.LoginService.Login(email, password)

			if err != nil {
				log.Debugf("%+v", err)
				log.Fatal(err.Error())
			}
		},
	}

	command.Flags().StringVarP(&email, "email", "e", "", "email (required)")
	command.Flags().StringVarP(&password, "password", "p", "", "password can be set at prompt as well (optional)")

	err := command.MarkFlagRequired("email")

	if err != nil {
		log.Fatal(err)
	}

	return command
}

func (l *Login) askForPassword() (string, error) {

	var input string

	fmt.Print("password: ")

	_, err := fmt.Scanln(&input)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return input, nil
}
