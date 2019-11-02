package fileentry

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestFileEntry__FileEntry__IsProcessable(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, path, []byte{}, 0744)
		entry := NewFileEntry(path)

		actual := entry.IsProcessable(fs)
		assert.True(t, actual)
	})

	t.Run("return false", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			entry := NewFileEntry(path)

			actual := entry.IsProcessable(fs)
			assert.False(t, actual)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, path+".lock", []byte("2147483647"), 0744)
			entry := NewFileEntry(path)

			actual := entry.IsProcessable(fs)
			assert.False(t, actual)
		})
	})
}

func TestFileEntry__FileEntry__Src(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		entry := NewFileEntry(path)

		actual := entry.Src()
		assert.Equal(t, actual, entry.rawPath)
	})
}

func TestFileEntry__FileEntry__Remove(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, path, []byte{}, 0744)
		entry := NewFileEntry(path)

		err := entry.Remove(fs)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			entry := NewFileEntry(path)

			err := entry.Remove(fs)
			assert.Error(t, err)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, path+".lock", []byte("2147483647"), 0744)
			entry := NewFileEntry(path)

			err := entry.Remove(fs)
			assert.Error(t, err)
		})
	})
}
