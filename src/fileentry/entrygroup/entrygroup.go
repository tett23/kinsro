package entrygroup

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/intdate"
)

// EntryGroup EntryGroup
type EntryGroup struct {
	Date    intdate.IntDate
	Entries []fileentry.FileEntry
}

// NewEntryGroup NewEntryGroup
func NewEntryGroup(date intdate.IntDate, entries []fileentry.FileEntry) (*EntryGroup, error) {
	ret := EntryGroup{
		Date:    date,
		Entries: entries,
	}

	return &ret, nil
}

// NewEntryGroupFromTSPath NewEntryGroupFromTSPath
func NewEntryGroupFromTSPath(fs afero.Fs, date intdate.IntDate, src string) (*EntryGroup, error) {
	if !strings.HasSuffix(src, ".ts") {
		return nil, errors.Errorf("Invalid TS path. path=%v", src)
	}

	excludeExt := src[0 : len(src)-3]

	matches, err := afero.Glob(fs, excludeExt+"*")
	if err != nil {
		return nil, errors.Wrapf(err, "Glob error. path=%v", excludeExt)
	}

	var entries []fileentry.FileEntry
	for i := range matches {
		entry, err := fileentry.NewFileEntry(matches[i])
		if err != nil {
			continue
		}
		if entry.Ext() == ".ts" {
			continue
		}

		entries = append(entries, *entry)
	}

	return NewEntryGroup(date, entries)
}

// IsEmpty IsEmpty
func (group EntryGroup) IsEmpty() bool {
	return len(group.Entries) == 0
}
