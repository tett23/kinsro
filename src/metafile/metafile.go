package metafile

// MetaFile MetaFile
type MetaFile struct {
	rawPath string
}

// NewMetaFile NewMetaFile
func NewMetaFile(path string) (*MetaFile, error) {
	ret := MetaFile{
		rawPath: path,
	}

	return &ret, nil
}
