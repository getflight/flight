package commands

import (
	"github.com/getflight/flight/helpers"
	"github.com/getflight/flight/http"
	"github.com/getflight/flight/service"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

type Root struct {
	DeploymentService *service.DeploymentService
	LoginService      *service.LoginService
	VersionService    *service.VersionService
	verbose           bool
	workPath          string
	apiUrl            string
}

func (r *Root) Execute() {

	cobra.OnInitialize(r.initConfig)

	rootCmd := &cobra.Command{
		Use:   "flight",
		Short: "Used to interact with the flight api",
		Long:  `Deploy infinitely scalable serverless GO apps. Complete documentation is available at https://getflight.io`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	rootCmd.PersistentFlags().BoolVarP(&r.verbose, "verbose", "v", false, "print verbose logs")
	rootCmd.PersistentFlags().StringVar(&r.workPath, "work-path", "", "path to store local data")
	rootCmd.PersistentFlags().StringVar(&r.apiUrl, "api-url", "", "configure a different api for flight to use when running commands")

	rootCmd.AddCommand(r.deployCommand())
	rootCmd.AddCommand(r.loginCommand())
	rootCmd.AddCommand(r.versionCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func (r *Root) initConfig() {

	if r.verbose {
		log.SetLevel(log.DebugLevel)
	}

	if r.workPath != "" {
		helpers.UserWorkPath = r.workPath
	}

	if r.apiUrl != "" {
		http.CustomApiUrl = r.apiUrl
	}
}

func (r *Root) deployCommand() *cobra.Command {
	deploy := &Deploy{
		DeploymentService: r.DeploymentService,
	}

	return deploy.command()
}

func (r *Root) loginCommand() *cobra.Command {
	login := &Login{
		LoginService: r.LoginService,
	}

	return login.command()
}

func (r *Root) versionCommand() *cobra.Command {
	version := &Version{
		VersionService: r.VersionService,
	}

	return version.command()
}
