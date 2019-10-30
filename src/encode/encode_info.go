package encode

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"github.com/tett23/kinsro/src/config"
)

// EncodeInfo EncodeInfo
type EncodeInfo struct {
	RawPath           string
	MetadataFileNames []string
	StoragePaths      []string
	Storage           string
}

var tsFilenameRegexp = regexp.MustCompile("/\\d{8}.+\\.ts$")

// NewEncodeInfo NewEncodeInfo
func NewEncodeInfo(conf *config.Config, fs afero.Fs, tsPath string) (*EncodeInfo, error) {
	if !tsFilenameRegexp.MatchString(tsPath) {
		return nil, errors.Errorf("format error")
	}

	ret := EncodeInfo{
		RawPath:      tsPath,
		StoragePaths: conf.StoragePaths,
	}

	metaFiles, err := ret.FetchMeta(fs)
	if err != nil {
		return nil, err
	}
	ret.MetadataFileNames = metaFiles

	return &ret, nil
}

// Base Base
func (info EncodeInfo) Base() string {
	return filepath.Base(info.RawPath)
}

// Dir Dir
func (info EncodeInfo) Dir() string {
	return filepath.Dir(info.RawPath)
}

// Prefix Prefix
func (info EncodeInfo) Prefix() string {
	base := info.Base()
	pos := strings.LastIndex(base, ".ts")

	return base[0:pos]
}

// MoveFiles MoveFiles
func (info EncodeInfo) MoveFiles() []string {
	for _, item := range info.MetadataFileNames {
		if strings.HasSuffix(item, ".mp4") {
			return info.MetadataFileNames
		}
	}

	return append(info.MetadataFileNames, info.MP4Name())
}

// MP4Name MP4Name
func (info EncodeInfo) MP4Name() string {
	return info.Prefix() + ".mp4"
}

// TSDest TSDest
func (info EncodeInfo) TSDest() string {
	return filepath.Join(info.Dir(), info.MP4Name())
}

// Move Move
func (info EncodeInfo) Move(statfs Statfs, fs afero.Fs) error {
	dir := info.Dir()
	files := info.MoveFiles()
	for i := range files {
		enc, err := NewEncodeFilePath(statfs, filepath.Join(dir, files[i]), info.StoragePaths)
		if err != nil {
			errors.Wrapf(err, "")
		}

		if err := fs.MkdirAll(enc.Dir(), 0755); err != nil {
			return err
		}

		if err := enc.Copy(fs); err != nil {
			return errors.Wrapf(err, "rename failed. src=%v dst=%v", enc.Src(), enc.Dest())
		}
	}

	return nil
}

// FetchMeta FetchMeta
func (info EncodeInfo) FetchMeta(fs afero.Fs) ([]string, error) {
	entries, err := afero.ReadDir(fs, info.Dir())
	if err != nil {
		return nil, errors.Wrapf(err, "ReadDir faired. path =%v", info.Dir())
	}

	prefix := info.Prefix()
	filtered := []string{}
	for i := range entries {
		if entries[i].IsDir() {
			continue
		}

		name := entries[i].Name()
		if strings.HasSuffix(name, ".ts") {
			continue
		}
		if !strings.HasPrefix(name, prefix) {
			continue
		}

		filtered = append(filtered, name)
	}

	return filtered, nil

}
