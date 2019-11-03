package writer

import (
	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/storages"
	"github.com/tett23/kinsro/src/syscalls"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// CreateSymlink CreateSymlink
func CreateSymlink(ss *syscalls.Syscalls, fs afero.Fs, vindexItem *vindexdata.VIndexItem, storage *storages.Storage) error {
	src := filepath.Join(storage.Path, vindexItem.Path())
	dst := filepath.Join(storage.Path, vindexItem.SymlinkPath())

	err := createSymlink(ss, fs, src, dst)
	if err != nil {
		return errors.Wrapf(err, "createSymlink failed. src=%v dst=%v", src, dst)
	}

	return nil
}

func createSymlink(ss *syscalls.Syscalls, fs afero.Fs, src, dst string) error {
	currentPwd, err := filepath.Abs(".")
	if err != nil {
		return errors.Wrapf(err, "filepath.Abs failed.")
	}

	base := filepath.Dir(dst)
	fs.MkdirAll(base, 0755)

	relSrc, err := filepath.Rel(base, src)
	if err != nil {
		return errors.Wrapf(err, "filepath.Rel failed. src=%v dst=%v", base, src)
	}

	ss.Chdir(base)
	defer ss.Chdir(currentPwd)

	fs.Remove(dst)
	err = ss.Symlink(relSrc, filepath.Base(dst))
	if err != nil {
		return errors.Wrapf(err, "os.Symlinkfailed. src=%v dst=%v", relSrc, filepath.Base(dst))
	}

	return nil
}
