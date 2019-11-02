package moventries

import (
	"os"
	"sort"
	"syscall"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/entrygroup"
)

// Statfs syscall.Statfs
type Statfs func(path string, stat *syscall.Statfs_t) error

type storageStat struct {
	storage string
	stat    syscall.Statfs_t
}

// MoveEntryGroup MoveEntryGroup
func MoveEntryGroup(fs afero.Fs, statfs Statfs, group *entrygroup.EntryGroup, storages []string) error {
	storage, err := mostSpacefulStorage(statfs, storages)
	if err != nil {
		return err
	}

	for i := range group.Entries {
		src := group.Entries[i]
		dest, err := fileentry.NewFileEntryFromEntry(storage, group.Date, &src)
		if err != nil {
			return err
		}

		if err := Copy(fs, &src, dest); err != nil {
			return err
		}
		if err := src.Remove(fs); err != nil {
			return err
		}
	}

	return nil
}

// Copy Copy
func Copy(fs afero.Fs, src, dest *fileentry.FileEntry) error {
	srcFile, err := fs.OpenFile(src.Src(), os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	return afero.WriteReader(fs, dest.Src(), srcFile)
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
