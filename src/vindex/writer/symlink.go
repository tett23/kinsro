package writer

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// CreateSymlink CreateSymlink
func CreateSymlink(conf *config.Config, vindexItem *vindexdata.VIndexItem) error {
	m := storageMap(conf.StoragePaths)
	symlinkDir, ok := m[vindexItem.Storage]
	if !ok {
		return errors.Errorf("%v", vindexItem.Storage)
	}

	symlinkName := vindexItem.SymlinkName()
	src := vindexItem.Filename
	dst := filepath.Join(symlinkDir, symlinkName)

	err := createSymlink(src, dst)
	if err != nil {
		return errors.Wrapf(err, "createSymlink failed. src=%v dst=%v", src, dst)
	}

	return nil
}

func createSymlink(src, dst string) error {
	currentPwd, err := filepath.Abs(".")
	if err != nil {
		return errors.Wrapf(err, "filepath.Abs failed.")
	}

	base := filepath.Dir(dst)
	os.MkdirAll(base, 0755)

	relSrc, err := filepath.Rel(base, src)
	if err != nil {
		return errors.Wrapf(err, "filepath.Rel failed. src=%v dst=%v", base, src)
	}

	os.Chdir(base)
	defer os.Chdir(currentPwd)

	os.Remove(dst)
	err = os.Symlink(relSrc, filepath.Base(dst))
	if err != nil {
		fmt.Println("os.Symlink", err)
		return errors.Wrapf(err, "os.Symlinkfailed. src=%v dst=%v", relSrc, filepath.Base(dst))
	}

	return nil
}

func storageMap(storagePaths []string) map[string]string {
	ret := map[string]string{}
	for i := range storagePaths {
		basename := filepath.Base(storagePaths[i])
		ret[basename] = filepath.Join(storagePaths[i], "symlinks")
	}

	return ret
}
