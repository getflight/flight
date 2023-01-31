package commands

import (
	"github.com/getflight/flight/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestLoginCommand(t *testing.T) {
	t.Run("command returns not nil command", func(t *testing.T) {
		// given
		login := Login{}

		// when
		command := login.command()

		// then
		assert.NotNil(t, command)
	})

	t.Run("command run calls login service when command is ran", func(t *testing.T) {
		// given
		loginServiceMock := &mocks.LoginServiceMock{}
		loginServiceMock.On("Login", mock.Anything, mock.Anything).Return(nil)

		login := Login{
			LoginService: loginServiceMock,
		}

		command := login.command()

		// when
		command.Run(command, []string{})

		// then
		loginServiceMock.AssertExpectations(t)
	})
}
