package reader

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestReader__NewAbstructReaderFromFile(t *testing.T) {
	vindexPath := "/tmp/vindex"

	t.Run("ok", func(t *testing.T) {
		t.Run("file exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, vindexPath, []byte{}, 0644)

			areader, err := NewAbstructReaderFromFile(fs, vindexPath)
			assert.NoError(t, err)
			assert.NotNil(t, areader)
		})

		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			areader, err := NewAbstructReaderFromFile(fs, vindexPath)
			assert.NoError(t, err)
			assert.NotNil(t, areader)
		})
	})
}

func TestReader__Close(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("file exists", func(t *testing.T) {
			areader := NewAbstructReaderFromBytes([]byte{})

			assert.NoError(t, areader.Close())
			assert.Error(t, areader.Close())
		})
	})
}
