package filesystem

import (
	"fmt"
	"github.com/spf13/afero"
	"io"
	"path/filepath"
)

func CopyDir(fs afero.Fs, from string, to string) error {
	fi, err := fs.Stat(from)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("%s is not a directory", from)
	}

	err = fs.MkdirAll(to, 0777) // before umask
	if err != nil {
		return err
	}

	entries, _ := afero.ReadDir(fs, from)
	for _, entry := range entries {
		fromFilename := filepath.Join(from, entry.Name())
		toFilename := filepath.Join(to, entry.Name())

		if entry.IsDir() {
			if err := CopyDir(fs, fromFilename, toFilename); err != nil {
				return err
			}
		} else {
			if err := CopyFile(fs, fromFilename, toFilename); err != nil {
				return err
			}
		}
	}

	return nil
}

func CopyFile(fs afero.Fs, from, to string) error {
	sf, err := fs.Open(from)
	if err != nil {
		return err
	}

	defer sf.Close()

	df, err := fs.Create(to)
	if err != nil {
		return err
	}

	defer df.Close()

	_, err = io.Copy(df, sf)
	if err != nil {
		return err
	}

	si, err := fs.Stat(from)
	if err != nil {
		err = fs.Chmod(to, si.Mode())

		if err != nil {
			return err
		}
	}

	return nil
}
