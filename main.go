package main

import (
	"github.com/getflight/flight/commands"
	"github.com/getflight/flight/formatters"
	"github.com/getflight/flight/helpers"
	"github.com/getflight/flight/http"
	"github.com/getflight/flight/service"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetFormatter(&formatters.TextFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)

	fileSystem := &helpers.FileSystem{}
	fileHelper := &helpers.FileHelper{FileSystem: fileSystem}
	tokenHelper := &helpers.TokenHelper{FileHelper: fileHelper}

	client := &http.Client{
		TokenHelper: tokenHelper,
	}

	versionService := &service.VersionService{}

	deploymentService := &service.DeploymentService{
		Client:      client,
		FileHelper:  fileHelper,
		TokenHelper: tokenHelper,
	}

	loginService := &service.LoginService{
		Client:      client,
		TokenHelper: tokenHelper,
	}

	root := &commands.Root{
		DeploymentService: deploymentService,
		LoginService:      loginService,
		VersionService:    versionService,
	}

	root.Execute()
}
