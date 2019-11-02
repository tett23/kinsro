package encode

import (
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
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

// Date Date
func (item EncodeFilePath) Date() uint64 {
	basename := item.Base()
	num, _ := strconv.Atoi(basename[0:8])

	return uint64(num)
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

// Copy Copy
func (item EncodeFilePath) Copy(fs afero.Fs) error {
	src, err := fs.OpenFile(item.Src(), os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer src.Close()

	return afero.WriteReader(fs, item.Dest(), src)
}

// ToVIndexItem ToVIndexItem
func (item EncodeFilePath) ToVIndexItem() (*vindexdata.VIndexItem, error) {
	if !item.isEncoded() {
		return nil, errors.Errorf("EncodeFilePath is not mp4. path=%v", item.Path)
	}

	storageName := filepath.Base(item.Storage)
	vindexItem, err := vindexdata.NewVIndexItem(storageName, item.Date(), item.Base())
	if err != nil {
		return nil, err
	}

	return vindexItem, nil
}

func (item EncodeFilePath) isEncoded() bool {
	return strings.HasSuffix(item.Path, ".mp4")
}
