package moveentries

import (
	"os"
	"syscall"

	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/entrygroup"
	"github.com/tett23/kinsro/src/storages"
)

type storageStat struct {
	storage string
	stat    syscall.Statfs_t
}

// MoveEntryGroup MoveEntryGroup
func MoveEntryGroup(fs afero.Fs, group *entrygroup.EntryGroup, storage *storages.Storage) error {
	for i := range group.Entries {
		src := group.Entries[i]
		dest, err := fileentry.NewFileEntryFromEntry(storage.Path, group.Date, &src)
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
