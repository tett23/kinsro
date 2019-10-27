package encode

import (
	"path/filepath"
	"sort"
	"syscall"

	"github.com/pkg/errors"
)

// EncodeFilePath EncodeFilePath
type EncodeFilePath struct {
	Path    string
	Storage string
}

// Statfs syscall.Statfs
type Statfs func(path string, stat *syscall.Statfs_t) error

// NewEncodeFilePath NewEncodeFilePath
func NewEncodeFilePath(statfs Statfs, path string, storages []string) (*EncodeFilePath, error) {
	storage, err := mostSpacefulStorage(statfs, storages)
	if err != nil {
		return nil, err
	}

	ret := EncodeFilePath{
		Path:    path,
		Storage: storage,
	}

	return &ret, nil
}

type storageStat struct {
	storage string
	stat    syscall.Statfs_t
}

func mostSpacefulStorage(statfs Statfs, storages []string) (string, error) {
	if len(storages) == 0 {
		return "", errors.Errorf("")
	}
	if len(storages) == 1 {
		return storages[0], nil
	}

	stats := make([]storageStat, len(storages))
	for i := range storages {
		stat := syscall.Statfs_t{}
		if err := statfs(storages[i], &stat); err != nil {
			return "", errors.Wrapf(err, "statfs failed. storage=%v", storages[i])
		}

		stats[i] = storageStat{
			storage: storages[i],
			stat:    stat,
		}
	}

	sort.Slice(stats, func(a, b int) bool {
		aFree := stats[a].stat.Bfree * uint64(stats[a].stat.Bsize)
		bFree := stats[b].stat.Bfree * uint64(stats[b].stat.Bsize)

		return aFree > bFree
	})

	return stats[0].storage, nil
}

// Src Src
func (item EncodeFilePath) Src() string {
	return item.Path
}

// Dest Dest
func (item EncodeFilePath) Dest() string {
	return filepath.Join(item.Dir(), item.Base())
}

// Dir Dir
func (item EncodeFilePath) Dir() string {
	basename := item.Base()
	y := basename[0:4]
	m := basename[4:6]
	d := basename[6:8]

	return filepath.Join(item.Storage, y, m, d)
}

// Base Base
func (item EncodeFilePath) Base() string {
	return filepath.Base(item.Path)
}
