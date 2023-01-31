package service

import (
	"github.com/getflight/flight/mocks"
	"github.com/getflight/flight/models"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestLoginService(t *testing.T) {
	t.Run("Login with client login error returns error", func(t *testing.T) {
		// given
		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", mock.Anything).Return(models.Token{}, errors.New("test error"))

		loginService := LoginService{
			Client: clientMock,
		}

		// when
		err := loginService.Login("test", "test")

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("Login with save token error returns error", func(t *testing.T) {
		// given
		email := "test@test.com"
		password := "test"

		login := models.Login{
			Email:    email,
			Password: password,
		}

		token := models.Token{
			Value: "123",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", login).Return(token, nil)

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("SaveToken", token.Value).Return(errors.New("test error"))

		loginService := LoginService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		err := loginService.Login(email, password)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("Login with get user error returns error", func(t *testing.T) {
		// given
		email := "test@test.com"
		password := "test"

		login := models.Login{
			Email:    email,
			Password: password,
		}

		token := models.Token{
			Value: "123",
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", login).Return(token, nil)
		clientMock.On("GetUser").Return(models.User{}, errors.New("test error"))

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("SaveToken", token.Value).Return(nil)

		loginService := LoginService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		err := loginService.Login(email, password)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("Login with get organisation error returns error", func(t *testing.T) {
		// given
		email := "test@test.com"
		password := "test"

		login := models.Login{
			Email:    email,
			Password: password,
		}

		token := models.Token{
			Value: "123",
		}

		user := models.User{
			ID:            "1",
			Email:         email,
			Organisations: []models.Organisation{},
			CreatedAt:     time.Time{},
			UpdatedAt:     time.Time{},
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", login).Return(token, nil)
		clientMock.On("GetUser").Return(user, nil)

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("SaveToken", token.Value).Return(nil)

		loginService := LoginService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		err := loginService.Login(email, password)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("Login with save organisation error returns error", func(t *testing.T) {
		// given
		email := "test@test.com"
		password := "test"

		login := models.Login{
			Email:    email,
			Password: password,
		}

		token := models.Token{
			Value: "123",
		}

		user := models.User{
			ID:    "1",
			Email: email,
			Organisations: []models.Organisation{
				{
					ID:   "1",
					Name: "org1",
				},
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", login).Return(token, nil)
		clientMock.On("GetUser").Return(user, nil)

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("SaveToken", token.Value).Return(nil)
		tokenHelperMock.On("SaveOrganisation", "1").Return(errors.New("test error"))

		loginService := LoginService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		err := loginService.Login(email, password)

		// then
		assert.NotNil(t, err)
		clientMock.AssertExpectations(t)
	})

	t.Run("Login with expected calls returns nil", func(t *testing.T) {
		// given
		email := "test@test.com"
		password := "test"

		login := models.Login{
			Email:    email,
			Password: password,
		}

		token := models.Token{
			Value: "123",
		}

		user := models.User{
			ID:    "1",
			Email: email,
			Organisations: []models.Organisation{
				{
					ID:   "1",
					Name: "org1",
				},
			},
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		}

		clientMock := &mocks.ClientMock{}
		clientMock.On("Login", login).Return(token, nil)
		clientMock.On("GetUser").Return(user, nil)

		tokenHelperMock := &mocks.TokenHelperMock{}
		tokenHelperMock.On("SaveToken", token.Value).Return(nil)
		tokenHelperMock.On("SaveOrganisation", "1").Return(nil)

		loginService := LoginService{
			Client:      clientMock,
			TokenHelper: tokenHelperMock,
		}

		// when
		err := loginService.Login(email, password)

		// then
		assert.Nil(t, err)
		clientMock.AssertExpectations(t)
	})
}
