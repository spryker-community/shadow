package config

import (
	qt "github.com/frankban/quicktest"
	"github.com/spf13/afero"
	"testing"
)

func TestReadLinksWithAbsentFile(t *testing.T) {
	c := qt.New(t)
	fs := afero.NewMemMapFs()

	result, err := ReadLinks(fs, "/path")

	c.Assert(result, qt.IsNil)
	c.Assert(err, qt.ErrorMatches, `(?s).*Unable to read file at ".*".*`)
}

func TestReadLinksWithInvalidYamlFile(t *testing.T) {
	c := qt.New(t)
	fs := afero.NewMemMapFs()

	data := `
*#invalid???"&"
`

	_ = afero.WriteFile(fs, "/path/file.yml", []byte(data), 0644)

	result, err := ReadLinks(fs, "/path/file.yml")

	c.Assert(result, qt.IsNil)
	c.Assert(err, qt.ErrorMatches, `(?s).*Invalid YAML file provided at ".*".*`)
}

func TestReadLinks(t *testing.T) {
	c := qt.New(t)
	fs := afero.NewMemMapFs()

	data := `
from: to
`

	_ = afero.WriteFile(fs, "/path/file.yml", []byte(data), 0644)

	result, err := ReadLinks(fs, "/path/file.yml")

	c.Assert(result, qt.IsNotNil)
	c.Assert(result, qt.DeepEquals, Links{"from": "to"})
	c.Assert(err, qt.IsNil)
}

func TestWriteLinksWithInvalidFilesystem(t *testing.T) {
	c := qt.New(t)
	fs := afero.NewReadOnlyFs(afero.NewMemMapFs())

	err := WriteLinks(fs, "/path/file.yml", Links{})

	c.Assert(err, qt.ErrorMatches, `(?s).*Unable to write file at ".*".*`)
}

func TestWriteLinks(t *testing.T) {
	c := qt.New(t)
	fs := afero.NewMemMapFs()

	err := WriteLinks(fs, "/path/file.yml", Links{})

	c.Assert(err, qt.IsNil)
}
