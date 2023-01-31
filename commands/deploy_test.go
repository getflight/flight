package commands

import (
	"github.com/getflight/flight/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestDeployCommand(t *testing.T) {
	t.Run("command returns not nil command", func(t *testing.T) {
		// given
		deploy := Deploy{}

		// when
		command := deploy.command()

		// then
		assert.NotNil(t, command)
	})

	t.Run("run command calls deployment service when command is ran", func(t *testing.T) {
		// given
		deploymentServiceMock := &mocks.DeploymentServiceMock{}
		deploymentServiceMock.On("Deploy", mock.Anything).Return(nil)

		deploy := Deploy{
			DeploymentService: deploymentServiceMock,
		}

		command := deploy.command()

		// when
		command.Run(command, []string{})

		// then
		deploymentServiceMock.AssertExpectations(t)
	})
}
