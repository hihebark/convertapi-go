package param

import (
	"fmt"
	"github.com/hihebark/convertapi-go/config"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ParamFile struct {
	ParamReader
	filePath string
}

func NewFile(name string, file *os.File, conf *config.Config) *ParamFile {
	paramReader := NewReader(name, file, filepath.Base(file.Name()), conf)
	return &ParamFile{*paramReader, file.Name()}
}

func NewFileReader(name string, reader io.Reader, conf *config.Config) *ParamFile {
	paramReader := NewReader("file", reader, name, conf)
	return &ParamFile{*paramReader, ""}
}

func NewPath(name string, path string, conf *config.Config) IParam {
	file, err := os.Open(path)
	if err != nil {
		return NewError(name, err)
	}
	return NewFile(name, file, conf)
}

func (this *ParamFile) Prepare() error {
	if this.filePath == "" {
		return this.ParamReader.Prepare()
	}
	file, err := os.Open(this.filePath)
	if err != nil {
		return err
	}
	this.reader = file
	return this.ParamReader.Prepare()
}

func (this *ParamFile) Values() ([]string, error) {
	err := this.Prepare()
	return this.values, err
}

func (this *ParamFile) String() string {
	return fmt.Sprintf("%s: %s -> %s", this.name, this.filePath, strings.Join(this.values, " "))
}
