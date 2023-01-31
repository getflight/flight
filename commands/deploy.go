package commands

import (
	"github.com/getflight/flight/formatters"
	"github.com/getflight/flight/service"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type Deploy struct {
	DeploymentService service.DeploymentServiceType
}

func (d *Deploy) command() *cobra.Command {
	var environment string

	command := &cobra.Command{
		Use:   "deploy",
		Short: "Deploy your code to flight's infrastructure",
		Long:  `Deploy will take your local code and make it available in flight's serverless infrastructure`,
		Run: func(cmd *cobra.Command, args []string) {
			log.SetFormatter(&formatters.TimestampFormatter{})

			err := d.DeploymentService.Deploy(environment)

			if err != nil {
				log.Debugf("%+v", err)
				log.Fatal(err.Error())
			}
		},
	}

	command.Flags().StringVarP(&environment, "environment", "e", "", "environment to deploy to (required)")

	err := command.MarkFlagRequired("environment")

	if err != nil {
		log.Fatal(err)
	}

	return command
}
