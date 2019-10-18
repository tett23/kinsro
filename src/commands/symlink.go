package commands

import (
	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/writer"
)

// Symlink Symlink
func Symlink(conf *config.Config, digest string) error {
	vindexItem, err := reader.FindByDigest(conf, digest)
	if err != nil {
		return errors.Wrapf(err, "reader.FindByDigest failed")
	}
	if vindexItem == nil {
		return errors.Errorf("digest not found. digest=%v", digest)
	}

	err = writer.CreateSymlink(conf, vindexItem)
	if err != nil {
		return errors.Wrapf(err, "writer.CreateSymlink failed")
	}

	return nil
}
