package commands

import (
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ListIndex ListIndex
func ListIndex(fs afero.Fs, indexPath string) (vindexdata.VIndex, error) {
	vindex, err := reader.ReadAll(fs, indexPath)
	if err != nil {
		return nil, err
	}

	return vindex, nil
}
