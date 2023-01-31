package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionService(t *testing.T) {
	t.Run("GetVersion returns value", func(t *testing.T) {
		// given
		versionService := VersionService{}

		// when
		version := versionService.GetVersion()

		// then
		assert.NotEmpty(t, version)
	})
}
