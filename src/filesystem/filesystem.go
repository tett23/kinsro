package filesystem

import (
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
)

var fs afero.Fs

func init() {
	config := config.GetConfig()
	if config.Environment == "test" {
		fs = afero.NewMemMapFs()
	} else {
		fs = afero.NewOsFs()
	}
}

// GetFs GetFs
func GetFs() afero.Fs {
	return fs
}
