package storages

import (
	"path/filepath"
	"syscall"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/syscalls"
)

// Storage Storage
type Storage struct {
	Path string
}

// NewStorage NewStorage
func NewStorage(path string) *Storage {
	ret := Storage{
		Path: path,
	}

	return &ret
}

// Name Name
func (storage Storage) Name() string {
	return filepath.Base(storage.Path)
}

// Free Free
func (storage Storage) Free(statfs syscalls.Statfs) (uint64, error) {
	stat := syscall.Statfs_t{}
	if err := statfs(storage.Path, &stat); err != nil {
		return 0, errors.Wrapf(err, "statfs failed. storage=%v", storage)
	}

	ret := stat.Bfree * uint64(stat.Bsize)

	return ret, nil
}
