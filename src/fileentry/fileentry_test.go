package fileentry

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/intdate"
)

func TestFileEntry__NewFileEntry(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		actual, err := NewFileEntry("/test.ts")
		assert.NoError(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("is not absolute path", func(t *testing.T) {
			actual, err := NewFileEntry("test.ts")
			assert.Error(t, err)
			assert.Nil(t, actual)
		})

		t.Run("is lock file", func(t *testing.T) {
			actual, err := NewFileEntry("/test.ts.lock")
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	})
}

func TestFileEntry__NewFileEntryFromEntry(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		date, _ := intdate.NewIntDate(20200101)
		entry, _ := NewFileEntry("/test.mp4")

		actual, err := NewFileEntryFromEntry("/media/video1", date, entry)
		assert.NoError(t, err)
		assert.NotNil(t, actual)
		assert.Equal(t, actual.Src(), "/media/video1/2020/01/01/test.mp4")
	})
}

func TestFileEntry__FileEntry__IsProcessable(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, path, []byte{}, 0744)
		entry, _ := NewFileEntry(path)

		actual := entry.IsProcessable(fs)
		assert.True(t, actual)
	})

	t.Run("return false", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			entry, _ := NewFileEntry(path)

			actual := entry.IsProcessable(fs)
			assert.False(t, actual)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, path+".lock", []byte("2147483647"), 0744)
			entry, _ := NewFileEntry(path)

			actual := entry.IsProcessable(fs)
			assert.False(t, actual)
		})
	})
}

func TestFileEntry__FileEntry__Src(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		entry, _ := NewFileEntry(path)

		actual := entry.Src()
		assert.Equal(t, actual, entry.rawPath)
	})
}

func TestFileEntry__FileEntry__Ext(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		entry, _ := NewFileEntry(path)

		actual := entry.Ext()
		assert.Equal(t, actual, ".ts")
	})
}

func TestFileEntry__FileEntry__Base(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		entry, _ := NewFileEntry(path)

		actual := entry.Base()
		assert.Equal(t, actual, "20191102_test.ts")
	})
}

func TestFileEntry__FileEntry__Remove(t *testing.T) {
	path := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, path, []byte{}, 0744)
		entry, _ := NewFileEntry(path)

		err := entry.Remove(fs)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			entry, _ := NewFileEntry(path)

			err := entry.Remove(fs)
			assert.Error(t, err)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, path+".lock", []byte("2147483647"), 0744)
			entry, _ := NewFileEntry(path)

			err := entry.Remove(fs)
			assert.Error(t, err)
		})
	})
}

func TestFileEntry__IsValidFileEntry(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ok := isValidFileEntry("/test.ts")
		assert.True(t, ok)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("is not absolute path", func(t *testing.T) {
			ok := isValidFileEntry("test.ts")
			assert.False(t, ok)
		})

		t.Run("is lock file", func(t *testing.T) {
			ok := isValidFileEntry("/test.ts.lock")
			assert.False(t, ok)
		})
	})
}
