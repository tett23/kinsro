package reader

import (
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ReadAll ReadAll
func ReadAll(vindexPath string) (vindexdata.VIndex, error) {
	fs := filesystem.GetFs()

	_, err := fs.Stat(vindexPath)
	if err != nil {
		return vindexdata.VIndex{}, err
	}

	bytes, err := afero.ReadFile(fs, vindexPath)
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
