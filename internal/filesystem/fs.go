package filesystem

import (
	"github.com/spf13/afero"
	"os"
	"path/filepath"
)

func NewOsFs() afero.Fs {
	return afero.NewOsFs()
}

func Glob(fs afero.Fs, pattern string) ([]string, error) {
	return afero.Glob(fs, pattern)
}

func Walk(fs afero.Fs, root string, walkFn filepath.WalkFunc) error {
	return afero.Walk(fs, root, walkFn)
}

func DirExists(fs afero.Fs, path string) (bool, error) {
	return afero.DirExists(fs, path)
}

func IsDir(fs afero.Fs, path string) (bool, error) {
	return afero.IsDir(fs, path)
}

func Exists(fs afero.Fs, path string) (bool, error) {
	return afero.Exists(fs, path)
}

func ReadFile(fs afero.Fs, path string) ([]byte, error) {
	return afero.ReadFile(fs, path)
}

func WriteFile(fs afero.Fs, path string, data []byte, perm os.FileMode) error {
	return afero.WriteFile(fs, path, data, perm)
}
