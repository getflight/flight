package helpers

import (
	"archive/zip"
	"fmt"
	"github.com/getflight/flight/models"
	"github.com/spf13/afero"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

const (
	flightWorkPath       = "/.flight/"
	buildWorkPath        = "build/"
	bootstrapFilename    = "bootstrap"
	executableFilename   = "main"
	organisationFilename = "organisation"
	tokenFilename        = "token"
	zipFilename          = "main.zip"
)

var (
	UserWorkPath = ""
)

type FileHelperType interface {
	Package(manifest models.Manifest) (string, error)
	ReadFile(filename string) (string, error)
	WriteFile(value string, filename string) error
}

type FileHelper struct {
	FileSystem FileSystemType
}

// Package bundles an executable into a zip file in order to prepare for the lambda deployment
func (h *FileHelper) Package(manifest models.Manifest) (string, error) {

	// Make sure the working directory exists
	err := h.prepareWrite()

	if err != nil {
		return "", errors.WithStack(err)
	}

	// Zip the executable
	err = h.writeZip(zipFilename, manifest)

	if err != nil {
		return "", errors.WithStack(err)
	}

	// Get the file handle
	content, err := h.ReadFile(buildWorkPath + zipFilename)

	if err != nil {
		return "", errors.WithStack(err)
	}

	return content, nil
}

func (h *FileHelper) ReadFile(filename string) (string, error) {
	path, err := h.getWorkPath(filename)

	if err != nil {
		return "", err
	}

	log.Debugf("reading from %v", path)

	content, err := h.FileSystem.ReadFile(path)

	return string(content), err
}

func (h *FileHelper) WriteFile(value string, filename string) error {
	err := h.prepareWrite()

	if err != nil {
		return errors.WithStack(err)
	}

	path, err := h.getWorkPath(filename)

	if err != nil {
		return err
	}

	log.Debugf("writing to %v %v", path, value)

	return h.FileSystem.WriteFile(path, []byte(value), 0644)
}

func (h *FileHelper) getWorkPath(filename string) (string, error) {
	var workPath string

	userHomeDir, err := h.FileSystem.UserHomeDir()

	if UserWorkPath != "" {
		workPath = UserWorkPath
	} else {
		workPath = userHomeDir
	}

	log.Debugf("default work path %v", userHomeDir)
	log.Debugf("provided work path %v", UserWorkPath)
	log.Debugf("used work path %v", workPath)

	return filepath.FromSlash(workPath + flightWorkPath + filename), err
}

func (h *FileHelper) prepareWrite() error {
	baseWorkPath, err := h.getWorkPath(buildWorkPath)

	if err != nil {
		return err
	}

	return h.FileSystem.MkdirAll(baseWorkPath, os.ModeDir)
}

func (h *FileHelper) openFile(filename string) (afero.File, error) {
	path, err := h.getWorkPath(filename)

	if err != nil {
		return nil, err
	}

	log.Debugf("opening file from %v", path)

	file, err := h.FileSystem.Open(path)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return file, nil
}

func (h *FileHelper) writeZip(zipFilename string, manifest models.Manifest) error {

	zipPath, err := h.getWorkPath(buildWorkPath + zipFilename)

	if err != nil {
		return errors.WithStack(err)
	}

	log.Debugf("zipping file to %s", zipPath)

	zipFile, err := h.FileSystem.Create(zipPath)

	defer func() {
		closeErr := zipFile.Close()
		if closeErr != nil {
			log.Errorf("failed to close zip file %v", err)
		}
	}()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = h.writeZipBootstrap(zipWriter)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, fmt.Sprintf("error while writing bootstrap")))
	}

	err = h.writeZipExecutable(zipWriter, manifest, executableFilename)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, fmt.Sprintf("error while writing executable")))
	}

	err = h.writeZipIncludes(zipWriter, manifest)

	if err != nil {
		return errors.WithStack(errors.Wrap(err, fmt.Sprintf("error while writing includes")))
	}

	return nil
}

func (h *FileHelper) writeZipExecutable(writer *zip.Writer, manifest models.Manifest, destination string) error {

	if manifest.Name == "" {
		return errors.WithStack(errors.New("name in manifest cannot be empty"))
	}

	exeWriter, err := writer.CreateHeader(&zip.FileHeader{
		CreatorVersion: 3 << 8,     // indicates Unix
		ExternalAttrs:  0777 << 16, // -rwxrwxrwx file permissions
		Name:           destination,
		Method:         zip.Deflate,
	})

	if err != nil {
		return errors.WithStack(err)
	}

	data, err := h.FileSystem.ReadFile(manifest.Name)

	if err != nil {
		return errors.WithStack(err)
	}

	_, err = exeWriter.Write(data)

	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (h *FileHelper) writeZipIncludes(writer *zip.Writer, manifest models.Manifest) error {
	if manifest.Files == nil {
		return nil
	}

	for _, include := range *manifest.Files {
		err := filepath.Walk(include,
			func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return errors.WithStack(err)
				}

				if !info.IsDir() {

					data, err := h.FileSystem.ReadFile(path)

					if err != nil {
						return errors.WithStack(err)
					}

					// rewrite path for unix specific path separators
					path = strings.ReplaceAll(path, "\\", "/")
					log.Debug(path)

					err = h.writeZipFile(writer, path, data)

					if err != nil {
						return errors.WithStack(err)
					}
				}

				return nil
			})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func (h *FileHelper) writeZipBootstrap(writer *zip.Writer) error {
	return h.writeZipFile(writer, bootstrapFilename, []byte(executableFilename))
}

func (h *FileHelper) writeZipFile(writer *zip.Writer, name string, content []byte) error {
	header := &zip.FileHeader{Name: name, Method: zip.Deflate}
	header.SetMode(0755)
	link, err := writer.CreateHeader(header)

	if err != nil {
		return errors.WithStack(err)
	}

	if _, err = link.Write(content); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
