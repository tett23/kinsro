package reader

import (
	"fmt"

	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
)

// ReadAll ReadAll
func ReadAll() (string, error) {
	config := config.GetConfig()
	fs := filesystem.GetFs()

	_, err := fs.Stat(config.VIndexPath)
	if err != nil {
		return "", err
	}

	bytes, err := afero.ReadFile(fs, config.VIndexPath)
	if err != nil {
		return "", err
	}
	fmt.Println(string(bytes))

	return string(bytes), nil
}
