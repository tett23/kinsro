package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/storages"
	"github.com/tett23/kinsro/src/syscalls"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/writer"
)

// SymlinkAll SymlinkAll
func SymlinkAll(conf *config.Config, ss *syscalls.Syscalls, fs afero.Fs) error {
	vindex, err := reader.ReadAll(fs, conf.VIndexPath)
	if err != nil {
		return errors.Cause(err)
	}

	storageItems := storages.NewStoragesFromPaths(conf.StoragePaths)

	for i := range vindex {
		storage, ok := storageItems.FindByName(vindex[i].Storage)
		if !ok {
			continue
		}

		err = writer.CreateSymlink(ss, fs, &vindex[i], storage)
		if err != nil {
			return errors.Cause(err)
		}
	}

	return nil
}
