package fileentry

import (
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/filelock"
)

// FileEntry FileEntry
type FileEntry struct {
	rawPath string
}

// NewFileEntry NewFileEntry
func NewFileEntry(path string) *FileEntry {
	ret := FileEntry{
		rawPath: path,
	}

	return &ret
}

// Src Src
func (entry FileEntry) Src() string {
	return entry.rawPath
}

// Remove Remove
func (entry FileEntry) Remove(fs afero.Fs) error {
	if !entry.isFree(fs) {
		return errors.Errorf("File does not processable. path=%v", entry.rawPath)
	}
	if err := fs.Remove(entry.rawPath); err != nil {
		return errors.Wrapf(err, "Remove failed. path=%v", entry.rawPath)
	}

	return nil
}

// IsProcessable IsProcessable
func (entry FileEntry) IsProcessable(fs afero.Fs) bool {
	return entry.isExists(fs) && entry.isFree(fs)
}

func (entry FileEntry) isExists(fs afero.Fs) bool {
	ok, err := afero.Exists(fs, entry.rawPath)

	return ok && err == nil
}

func (entry FileEntry) isFree(fs afero.Fs) bool {
	ok, err := filelock.IsFree(fs, entry.rawPath)

	return ok && err == nil
}
