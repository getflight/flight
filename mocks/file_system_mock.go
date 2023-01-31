package mocks

import (
	"archive/zip"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/mock"
	"io"
	"io/fs"
)

type FileSystemMock struct {
	mock.Mock
}

func (m *FileSystemMock) Create(name string) (afero.File, error) {
	args := m.Called(name)

	return args.Get(0).(afero.File), args.Error(1)
}

func (m *FileSystemMock) MkdirAll(path string, perm fs.FileMode) error {
	args := m.Called(path, perm)

	return args.Error(0)
}

func (m *FileSystemMock) NewWriter(w io.Writer) *zip.Writer {
	args := m.Called(w)

	return args.Get(0).(*zip.Writer)
}

func (m *FileSystemMock) Open(name string) (afero.File, error) {
	args := m.Called(name)

	return args.Get(0).(afero.File), args.Error(1)
}

func (m *FileSystemMock) ReadFile(filename string) ([]byte, error) {
	args := m.Called(filename)

	return args.Get(0).([]byte), args.Error(1)
}

func (m *FileSystemMock) UserHomeDir() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}

func (m *FileSystemMock) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	args := m.Called(filename, data, perm)

	return args.Error(0)
}
