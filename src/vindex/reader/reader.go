package reader

import (
	"crypto/md5"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// ReadAll ReadAll
func ReadAll(vindexPath string) (vindexdata.VIndex, error) {
	fs := filesystem.GetFs()

	_, err := fs.Stat(vindexPath)
	if err != nil {
		return vindexdata.VIndex{}, err
	}

	bytes, err := afero.ReadFile(fs, vindexPath)
	if err != nil {
		return vindexdata.VIndex{}, err
	}
	//
	vindex, err := vindexdata.NewVIndexFromBinary(bytes)
	if err != nil {
		return vindexdata.VIndex{}, err
	}

	return vindex, nil
}

// FindByFilename FindByFilename
func FindByFilename(conf *config.Config, filename string) (*vindexdata.VIndexItem, error) {
	fs := filesystem.GetFs()

	f, err := fs.OpenFile(conf.VIndexPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil, errors.Wrap(err, "OpenFile failed.")
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, errors.Wrap(err, "Stat failed")
	}

	digest := md5.Sum([]byte(filename))
	rowLen := vindexdata.RowLength()
	data := make([]byte, rowLen)
	fileSize := stat.Size()
	fileLen := fileSize / rowLen
	for i := int64(0); i < fileLen; i++ {
		_, err = f.ReadAt(data, i*rowLen)
		if err != nil {
			return nil, errors.Wrap(err, "ReadAt failed")
		}

		vindexItem, err := vindexdata.NewBinaryIndexItemFromBinary(data)
		if err != nil {
			return nil, errors.Wrap(err, "NewBinaryIndexItemFromBinary failed")
		}

		if vindexItem.Digest == digest {
			return vindexItem, nil
		}
	}

	return nil, nil
}
