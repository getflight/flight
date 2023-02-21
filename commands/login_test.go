package commands

import (
	"github.com/getflight/flight/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"io/ioutil"
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

		input, err := ioutil.TempFile("", "")

		if err != nil {
			t.Fatal(err)
		}
		defer input.Close()

		_, err = io.WriteString(input, "test")

		if err != nil {
			t.Fatal(err)
		}

		_, err = input.Seek(0, io.SeekStart)

		if err != nil {
			t.Fatal(err)
		}

		login := Login{
			LoginService: loginServiceMock,
			Input:        input,
		}

		command := login.command()

		// when
		command.Run(command, []string{})

		// then
		loginServiceMock.AssertExpectations(t)
	})
}
