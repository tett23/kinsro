package commands

import (
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ListIndex ListIndex
func ListIndex(indexPath string) (vindexdata.VIndex, error) {
	vindex, err := reader.ReadAll(indexPath)
	if err != nil {
		return nil, err
	}

	return vindex, nil
}
