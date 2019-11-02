package metafile

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/fileentry"
)

// MetaFile MetaFile
type MetaFile struct {
	*fileentry.FileEntry
}

// NewMetaFile NewMetaFile
func NewMetaFile(path string) (*MetaFile, error) {
	if !isValidMetaFile(path) {
		return nil, errors.Errorf("Invalid path. path=%v", path)
	}

	entry, err := fileentry.NewFileEntry(path)
	if err != nil {
		return nil, err
	}

	ret := MetaFile{
		FileEntry: entry,
	}

	return &ret, nil
}

func isValidMetaFile(path string) bool {
	if strings.HasSuffix(path, ".ts") {
		return false
	}
	if strings.HasSuffix(path, ".mp4") {
		return false
	}

	return true
}
