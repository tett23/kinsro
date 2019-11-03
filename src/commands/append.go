package commands

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

// AppendToIndex AppendToIndex
func AppendToIndex(conf *config.Config, fs afero.Fs, vindexItem *vindexdata.VIndexItem) error {
	_, ok, err := reader.FindByDigest(conf, fs, vindexItem.HexDigest())
	if err != nil {
		return err
	}
	if ok {
		return nil
	}

	ch := writer.Append(conf, vindexItem)
	ok = <-ch
	if !ok {
		return errors.Errorf("Append failed. vindex=%+v", vindexItem)
	}

	return nil
}
