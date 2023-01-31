package commands

import (
	"github.com/getflight/flight/service"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

type Version struct {
	VersionService *service.VersionService
}

func (v *Version) command() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `All software has versions. This is flight's`,
		Run: func(cmd *cobra.Command, args []string) {
			version := v.VersionService.GetVersion()
			log.Info(version)
		},
	}
}
