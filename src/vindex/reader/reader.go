package reader

import (
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ReadAll ReadAll
func ReadAll() (vindexdata.VIndex, error) {
	config := config.GetConfig()
	fs := filesystem.GetFs()

	_, err := fs.Stat(config.VIndexPath)
	if err != nil {
		return vindexdata.VIndex{}, err
	}

	bytes, err := afero.ReadFile(fs, config.VIndexPath)
	if err != nil {
		return vindexdata.VIndex{}, err
	}
	//
	vindex, err := vindexdata.NewVIndexFromBinary(bytes)
	if err != nil {
		return vindexdata.VIndex{}, err
	}

	return vindex, nil
}
