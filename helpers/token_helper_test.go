package helpers

import (
	"github.com/getflight/flight/mocks"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTokenHelper(t *testing.T) {
	t.Run("TokenExists with token returns true", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "token").Return("123", nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		result := tokenHelper.TokenExists()

		// then
		assert.True(t, result)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("TokenExists with empty token returns false", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "token").Return("", nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		result := tokenHelper.TokenExists()

		// then
		assert.False(t, result)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("TokenExists with error returns false", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "token").Return("", errors.New("error"))

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		result := tokenHelper.TokenExists()

		// then
		assert.False(t, result)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("SaveToken with success writes and returns nil", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("WriteFile", "123", "token").Return(nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		err := tokenHelper.SaveToken("123")

		// then
		assert.Nil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("SaveToken with error returns error", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("WriteFile", "123", "token").Return(errors.New("error"))

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		err := tokenHelper.SaveToken("123")

		// then
		assert.NotNil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("GetToken returns token", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "token").Return("123", nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		token, err := tokenHelper.GetToken()

		// then
		assert.Equal(t, "123", token)
		assert.Nil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("GetToken with error returns error", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "token").Return("", errors.New("error"))

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		token, err := tokenHelper.GetToken()

		// then
		assert.Equal(t, "", token)
		assert.NotNil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("SaveOrganisation with success returns nil", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("WriteFile", "1", "organisation").Return(nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		err := tokenHelper.SaveOrganisation("1")

		// then
		assert.Nil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("SaveOrganisation with error returns error", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("WriteFile", "1", "organisation").Return(errors.New("error"))

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		err := tokenHelper.SaveOrganisation("1")

		// then
		assert.NotNil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("GetOrganisation with success returns nil", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "organisation").Return("1", nil)

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		organisationId, err := tokenHelper.GetOrganisation()

		// then
		assert.Equal(t, "1", organisationId)
		assert.Nil(t, err)
		fileHelperMock.AssertExpectations(t)
	})

	t.Run("GetOrganisation with error returns error", func(t *testing.T) {
		// given
		fileHelperMock := &mocks.FileHelperMock{}
		fileHelperMock.On("ReadFile", "organisation").Return("", errors.New("error"))

		tokenHelper := TokenHelper{FileHelper: fileHelperMock}

		// when
		organisationId, err := tokenHelper.GetOrganisation()

		// then
		assert.Equal(t, "", organisationId)
		assert.Nil(t, err)
		fileHelperMock.AssertExpectations(t)
	})
}
