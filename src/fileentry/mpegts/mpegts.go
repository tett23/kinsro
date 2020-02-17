package mpegts

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/entrygroup"
	"github.com/tett23/kinsro/src/intdate"
	"github.com/tett23/kinsro/src/storages"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

// MpegTS MpegTS
type MpegTS struct {
	*fileentry.FileEntry
}

var tsFilenameRegexp = regexp.MustCompile("/\\d{8}.+\\.ts$")

// NewMpegTS NewMpegTS
func NewMpegTS(path string) (*MpegTS, error) {
	entry, err := fileentry.NewFileEntry(path)
	if err != nil {
		return nil, err
	}
	if !isMpegTSPath(path) {
		return nil, fmt.Errorf("Invalid TS path. %v", path)
	}

	ret := MpegTS{
		FileEntry: entry,
	}

	return &ret, nil
}

// Dest Dest
func (ts MpegTS) Dest() string {
	src := ts.Src()

	return src[0:len(src)-3] + ".mp4"
}

// LogPath LogPath
func (ts MpegTS) LogPath() string {
	src := ts.Src()

	return src[0:len(src)-3] + ".log"
}

// ToEntryGroup ToEntryGroup
func (ts MpegTS) ToEntryGroup(fs afero.Fs) (*entrygroup.EntryGroup, error) {
	date, err := ts.ToIntDate()
	if err != nil {
		return nil, err
	}

	return entrygroup.NewEntryGroupFromTSPath(fs, date, ts.Src())
}

// ToIntDate ToIntDate
func (ts MpegTS) ToIntDate() (intdate.IntDate, error) {
	base := ts.Base()
	if len(base) < 8 {
		return -1, errors.Errorf("Atoi failed. path=%v", base)
	}

	date, err := strconv.Atoi(base[0:8])
	if date == 0 || err != nil {
		return -1, errors.Wrapf(err, "Atoi failed. path=%v", base)
	}

	return intdate.NewIntDate(date)
}

// ToVIndexItem ToVIndexItem
func (ts MpegTS) ToVIndexItem(storage *storages.Storage) (*vindexdata.VIndexItem, error) {
	date, err := ts.ToIntDate()
	if err != nil {
		return nil, err
	}

	mp4 := filepath.Base(ts.Dest())

	return vindexdata.NewVIndexItem(storage.Name(), int(date), mp4)
}

func isMpegTSPath(path string) bool {
	return tsFilenameRegexp.MatchString(path)
}
