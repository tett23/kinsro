package filesystem

import (
	"os"

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
