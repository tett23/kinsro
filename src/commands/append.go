package commands

import (
	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

// Append Append
func Append(conf *config.Config, vindexItem *vindexdata.VIndexItem) error {
	record, err := reader.FindByFilename(conf, vindexItem.Filename)
	if err != nil {
		return errors.Errorf("FindByFilename failed. vindex=%+v", vindexItem)
	}
	if record != nil {
		return nil
	}

	ch := writer.Append(conf, vindexItem)
	ok := <-ch
	if !ok {
		return errors.Errorf("Append failed. vindex=%+v", vindexItem)
	}

	return nil
}
