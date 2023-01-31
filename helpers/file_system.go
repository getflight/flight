package helpers

import (
	"archive/zip"
	"github.com/spf13/afero"
	"io"
	"io/fs"
	"io/ioutil"
	"os"
)

type FileSystemType interface {
	Create(name string) (afero.File, error)
	MkdirAll(path string, perm fs.FileMode) error
	NewWriter(w io.Writer) *zip.Writer
	Open(name string) (afero.File, error)
	ReadFile(filename string) ([]byte, error)
	UserHomeDir() (string, error)
	WriteFile(filename string, data []byte, perm fs.FileMode) error
}

type FileSystem struct {
}

func (s *FileSystem) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (s *FileSystem) MkdirAll(path string, perm fs.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (s *FileSystem) Open(name string) (afero.File, error) {
	return os.Open(name)
}

func (s *FileSystem) Create(name string) (afero.File, error) {
	return os.Create(name)
}

func (s *FileSystem) WriteFile(filename string, data []byte, perm fs.FileMode) error {
	return ioutil.WriteFile(filename, data, perm)
}

func (s *FileSystem) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (s *FileSystem) NewWriter(w io.Writer) *zip.Writer {
	return zip.NewWriter(w)
}
