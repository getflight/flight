package helpers

import (
	"github.com/getflight/flight/mocks"
	"github.com/getflight/flight/models"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func TestFileHelper(t *testing.T) {
	t.Run("getWorkPath returns formatted path", func(t *testing.T) {
		// given
		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}
		filename := "test"

		// when
		workPath, err := fileHelper.getWorkPath(filename)

		// then
		assert.Nil(t, err)
		assert.Equal(t, filepath.Join("home", ".flight", "test"), workPath)
		fileSystemMock.AssertExpectations(t)
	})

	t.Run("getWorkPath with userWorkPath returns formatted path", func(t *testing.T) {
		// given
		UserWorkPath = "work"
		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}
		filename := "test"

		// when
		workPath, err := fileHelper.getWorkPath(filename)

		// then
		assert.Nil(t, err)
		assert.Equal(t, filepath.Join("work", ".flight", "test"), workPath)
		fileSystemMock.AssertExpectations(t)

		// revert
		UserWorkPath = ""
	})

	t.Run("prepareWrite creates work directory", func(t *testing.T) {
		// given
		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)
		fileSystemMock.On("MkdirAll", filepath.Join("home", ".flight", "build"), os.ModeDir).Return(nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}

		// when
		err := fileHelper.prepareWrite()

		// then
		assert.Nil(t, err)
		fileSystemMock.AssertExpectations(t)
	})

	t.Run("WriteFile writes value to file", func(t *testing.T) {
		// given
		value := "test"
		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)
		fileSystemMock.On("MkdirAll", filepath.Join("home", ".flight", "build"), os.ModeDir).Return(nil)
		fileSystemMock.On("WriteFile", filepath.Join("home", ".flight", "token"), []byte(value), fs.FileMode(0644)).Return(nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}

		// when
		err := fileHelper.WriteFile(value, "token")

		// then
		assert.Nil(t, err)
		fileSystemMock.AssertExpectations(t)
	})

	t.Run("ReadFile reads value from file", func(t *testing.T) {
		// given
		value := "test"
		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)
		fileSystemMock.On("ReadFile", filepath.Join("home", ".flight", "token")).Return([]byte(value), nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}

		// when
		result, err := fileHelper.ReadFile("token")

		// then
		assert.Nil(t, err)
		assert.Equal(t, value, result)
		fileSystemMock.AssertExpectations(t)
	})

	t.Run("openFile returns file", func(t *testing.T) {
		// given
		memFs := new(afero.MemMapFs)
		f, err := afero.TempFile(memFs, "", "test")

		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)
		fileSystemMock.On("Open", filepath.Join("home", ".flight", "token")).Return(f, nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}

		// when
		_, err = fileHelper.openFile("token")

		// then
		assert.Nil(t, err)
		fileSystemMock.AssertExpectations(t)
	})

	t.Run("writeZip writes folder with correct files", func(t *testing.T) {
		// given
		memFs := new(afero.MemMapFs)
		f, err := afero.TempFile(memFs, "", "test")

		fileSystemMock := &mocks.FileSystemMock{}
		fileSystemMock.On("UserHomeDir").Return("home", nil)
		fileSystemMock.On("Create", filepath.Join("home", ".flight", "build", "main")).Return(f, nil)
		fileSystemMock.On("ReadFile", "test-name").Return([]byte{}, nil)

		fileHelper := FileHelper{FileSystem: fileSystemMock}

		manifest := models.Manifest{
			Name:         "test-name",
			Files:        nil,
			Trigger:      "",
			Environments: nil,
		}

		// when
		err = fileHelper.writeZip("main", manifest)

		// then
		assert.Nil(t, err)
		fileSystemMock.AssertExpectations(t)
	})
}
