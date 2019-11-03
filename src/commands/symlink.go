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

// Symlink Symlink
func Symlink(conf *config.Config, ss *syscalls.Syscalls, fs afero.Fs, digest string) error {
	vindexItem, ok, err := reader.FindByDigest(conf, fs, digest)
	if err != nil {
		return errors.Wrapf(err, "reader.FindByDigest failed")
	}
	if !ok {
		return errors.Errorf("digest not found. digest=%v", digest)
	}
	if vindexItem == nil {
		return errors.Errorf("digest not found. digest=%v", digest)
	}

	storage, ok := storages.NewStoragesFromPaths(conf.StoragePaths).FindByName(vindexItem.Storage)
	if !ok {
		return errors.Errorf("storage not found. storage=%v", vindexItem.Storage)
	}

	err = writer.CreateSymlink(ss, fs, vindexItem, storage)
	if err != nil {
		return errors.Wrapf(err, "writer.CreateSymlink failed")
	}

	return nil
}
