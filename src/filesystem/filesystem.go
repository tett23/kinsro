package filesystem

import (
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

var fs afero.Fs
var rawFs afero.Fs

func init() {
	goEnv := os.Getenv("GO_ENV")
	if goEnv == "test" {
		fs = afero.NewMemMapFs()
		rawFs = afero.NewOsFs()
	} else {
		fs = afero.NewOsFs()
		rawFs = fs
	}
}

// GetFs GetFs
func GetFs() afero.Fs {
	return fs
}

// GetRawFs GetRawFs
func GetRawFs() afero.Fs {
	return rawFs
}

// ResetTestFs ResetTestFs
func ResetTestFs() {
	fs = afero.NewMemMapFs()
}

// CopyDotfile CopyDotfile
func CopyDotfile() {
	rawFs := GetRawFs()

	dotenvPath := filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "tett23", "kinsro", ".env.test")
	content, _ := afero.ReadFile(rawFs, dotenvPath)

	memFs := GetFs()
	afero.WriteFile(memFs, "/.env", content, 0644)
}
