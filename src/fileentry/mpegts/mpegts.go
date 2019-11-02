package mpegts

import (
	"fmt"
	"regexp"

	"github.com/tett23/kinsro/src/fileentry"
)

// MpegTS MpegTS
type MpegTS struct {
	*fileentry.FileEntry
}

var tsFilenameRegexp = regexp.MustCompile("/\\d{8}.+\\.ts$")

// NewMpegTS NewMpegTS
func NewMpegTS(path string) (*MpegTS, error) {
	if !isMpegTSPath(path) {
		return nil, fmt.Errorf("Invalid TS path. %v", path)
	}

	ret := MpegTS{
		FileEntry: fileentry.NewFileEntry(path),
	}

	return &ret, nil
}

// Dest Dest
func (ts MpegTS) Dest() string {
	src := ts.Src()

	return src[0:len(src)-3] + ".mp4"
}

func isMpegTSPath(path string) bool {
	return tsFilenameRegexp.MatchString(path)
}
