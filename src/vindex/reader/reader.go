package reader

import (
	"io"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ReadAll ReadAll
func ReadAll(fs afero.Fs, vindexPath string) (vindexdata.VIndex, error) {
	bytes, err := afero.ReadFile(fs, vindexPath)
	if err != nil {
		return nil, err
	}

	vindex, err := vindexdata.NewVIndexFromBinary(bytes)
	if err != nil {
		return nil, err
	}

	return vindex, nil
}

// FindByFullPath FindByFullPath
func FindByFullPath(conf *config.Config, fs afero.Fs, path string) (*vindexdata.VIndexItem, bool, error) {
	vindexItem, err := vindexdata.ParseFullFilepath(conf.StoragePaths, path)
	if err != nil {
		return nil, false, err
	}

	return FindByDigest(conf, fs, vindexItem.HexDigest())
}

// FindByDigest FindByDigest
func FindByDigest(conf *config.Config, fs afero.Fs, digest string) (*vindexdata.VIndexItem, bool, error) {
	rr, err := NewAbstructReaderFromFile(fs, conf.VIndexPath)
	if err != nil {
		return nil, false, errors.Wrap(err, "OpenFile failed.")
	}
	defer rr.Close()

	rowLen := vindexdata.RowLength()
	i := int64(0)
	data := make([]byte, rowLen)
	for true {
		_, err = rr.ReadAt(data, i*rowLen)
		if err == io.EOF {
			return nil, false, nil
		}
		if err != nil {
			return nil, false, errors.Wrap(err, "ReadAt failed")
		}
		i++

		vindexItem, err := vindexdata.NewBinaryIndexItemFromBinary(data)
		if err != nil {
			return nil, false, err
		}

		if vindexItem.HexDigest() == digest {
			return vindexItem, true, nil
		}
	}

	return nil, false, nil
}
