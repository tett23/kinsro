package commands

import (
	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/writer"
)

// SymlinkAll SymlinkAll
func SymlinkAll(conf *config.Config) error {
	vindex, err := reader.ReadAll(conf.VIndexPath)
	if err != nil {
		return errors.Cause(err)
	}

	for i := range vindex {
		err = writer.CreateSymlink(conf, &vindex[i])
		if err != nil {
			return errors.Cause(err)
		}
	}

	return nil
}
