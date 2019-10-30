package commands

import (
	"fmt"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/filesystem"
)

// EncodeTS EncodeTS
func EncodeTS(conf *config.Config, tsPath string) error {
	fs := filesystem.GetFs()
	info, err := encode.NewEncodeInfo(conf, fs, tsPath)
	if err != nil {
		return errors.Wrapf(err, "initialize error. ts=%v", tsPath)
	}

	err = encode.Encode(conf, fs, info)
	if err != nil {
		return err
	}

	fileInfo, err := encode.NewEncodeFilePath(syscall.Statfs, info.TSDest(), conf.StoragePaths)
	if err != nil {
		return err
	}

	vindexItem, err := fileInfo.ToVIndexItem()
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n%v\n", vindexItem, vindexItem.HexDigest())

	if err := AppendToIndex(conf, vindexItem); err != nil {
		return err
	}

	if err := Symlink(conf, vindexItem.HexDigest()); err != nil {
		return err
	}

	return nil
}
