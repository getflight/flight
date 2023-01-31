package commands

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCommand(t *testing.T) {
	t.Run("Execute does not fail", func(t *testing.T) {
		// given
		root := &Root{}

		// when
		root.Execute()

		// then
		assert.NotNil(t, root)
	})
}
