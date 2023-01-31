package context

import (
	"github.com/getflight/flight/models"
	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

type ConfigurationType interface {
	Init() error
	GetManifest() (models.Manifest, error)
}

type Configuration struct {
}

func (c *Configuration) Init() error {
	viper.SetConfigName("flight")
	viper.AddConfigPath(".")

	log.Info("reading configuration")

	err := viper.ReadInConfig()

	if err != nil {
		return err
	}

	return nil
}

func (c *Configuration) GetManifest() (models.Manifest, error) {
	manifest := models.Manifest{}
	err := viper.Unmarshal(&manifest)

	if err != nil {
		return manifest, errors.WithStack(err)
	}

	return manifest, nil
}
