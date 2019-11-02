package reader

import (
	"bytes"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

// AbstructReader AbstructReader
type AbstructReader interface {
	io.ReaderAt
	io.Closer
}

// NewAbstructReaderFromFile NewAbstructReaderFromFile
func NewAbstructReaderFromFile(fs afero.Fs, vindexPath string) (AbstructReader, error) {
	file, err := fs.OpenFile(vindexPath, os.O_RDONLY, 0644)
	if err != nil {
		switch t := err.(type) {
		case *os.PathError:
			if t.Err == afero.ErrFileNotFound {
				return NewAbstructReaderFromBytes([]byte{}), nil
			}
		default:
			return nil, errors.Wrap(err, "OpenFile failed.")
		}
	}

	var ret AbstructReader = file

	return ret, nil
}

// BytesBuffer BytesBuffer
type BytesBuffer struct {
	*bytes.Reader
	isClosed bool
}

// Close Close
func (buf *BytesBuffer) Close() error {
	if buf.isClosed {
		return os.ErrClosed
	}

	buf.isClosed = true

	return nil
}

// NewAbstructReaderFromBytes NewAbstructReaderFromBytes
func NewAbstructReaderFromBytes(data []byte) AbstructReader {
	reader := bytes.NewReader(data)
	ret := BytesBuffer{Reader: reader, isClosed: false}

	return &ret
}
