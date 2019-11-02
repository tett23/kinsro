package metafile

import "github.com/tett23/kinsro/src/fileentry"

// MetaFile MetaFile
type MetaFile struct {
	*fileentry.FileEntry
}

// NewMetaFile NewMetaFile
func NewMetaFile(path string) *MetaFile {
	ret := MetaFile{
		FileEntry: fileentry.NewFileEntry(path),
	}

	return &ret
}
