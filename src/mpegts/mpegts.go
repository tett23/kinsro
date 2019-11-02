package mpegts

import (
	"fmt"
	"regexp"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/filelock"
)

// MpegTS MpegTS
type MpegTS struct {
	RawPath string
}

var tsFilenameRegexp = regexp.MustCompile("/\\d{8}.+\\.ts$")

// NewMpegTS NewMpegTS
func NewMpegTS(path string) (*MpegTS, error) {
	if !isMpegTSPath(path) {
		return nil, fmt.Errorf("Invalid TS path. %v", path)
	}

	ret := MpegTS{
		RawPath: path,
	}

	return &ret, nil
}

// IsEncodable IsEncodable
func (ts MpegTS) IsEncodable(fs afero.Fs) bool {
	return ts.isExists(fs) && ts.isFree(fs)
}

// Dest Dest
func (ts MpegTS) Dest() string {
	return ts.RawPath[0:len(ts.RawPath)-3] + ".mp4"
}

// Remove Remove
func (ts MpegTS) Remove(fs afero.Fs) error {
	if !ts.isFree(fs) {
		return errors.Errorf("TS locked. path=%v", ts.RawPath)
	}
	if err := fs.Remove(ts.RawPath); err != nil {
		return errors.Wrapf(err, "Remove failed. path=%v", ts.RawPath)
	}

	return nil
}

func (ts MpegTS) isExists(fs afero.Fs) bool {
	ok, err := afero.Exists(fs, ts.RawPath)

	return ok && err == nil
}

func (ts MpegTS) isFree(fs afero.Fs) bool {
	ok, err := filelock.IsFree(fs, ts.RawPath)

	return ok && err == nil
}

func isMpegTSPath(path string) bool {
	return tsFilenameRegexp.MatchString(path)
}
