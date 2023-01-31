package context

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfiguration(t *testing.T) {
	t.Run("Init returns config file not found error", func(t *testing.T) {
		// given
		configuration := Configuration{}

		// when
		err := configuration.Init()

		// then
		assert.NotNil(t, err)
		assert.IsType(t, viper.ConfigFileNotFoundError{}, err)
	})
}
