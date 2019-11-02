package mpegts

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/metafile"
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

// MetaFiles MetaFiles
func (ts MpegTS) MetaFiles(fs afero.Fs) ([]metafile.MetaFile, error) {
	src := ts.Src()
	excludeExt := src[0 : len(src)-3]

	var ret []metafile.MetaFile
	matches, err := afero.Glob(fs, excludeExt+"*")
	if err != nil {
		return ret, errors.Wrapf(err, "Glob error. path=%v", excludeExt)
	}

	for i := range matches {
		entry, err := metafile.NewMetaFile(matches[i])
		if err != nil {
			continue
		}

		ret = append(ret, *entry)
	}

	return ret, nil
}

func isMpegTSPath(path string) bool {
	return tsFilenameRegexp.MatchString(path)
}
