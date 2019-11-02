package entrygroup

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/intdate"
)

func TestEntryGroup__NewEntryGroupFromTSPath(t *testing.T) {
	path := "/media/video_tmp/20191102_test.ts"
	createFiles := func(fs afero.Fs, files []string) {
		for i := range files {
			afero.WriteFile(fs, files[i], []byte{}, 0644)
		}
	}

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		date, _ := intdate.NewIntDate(20191102)
		createFiles(fs, []string{
			path,
			"/media/video_tmp/20191102_test.ts.lock",
			"/media/video_tmp/20191102_test.mp4",
			"/media/video_tmp/20191102_test.json",
			"/media/video_tmp/20191102_test.log",
			"/media/video_tmp/20191103_test.log",
		})

		group, err := NewEntryGroupFromTSPath(fs, date, path)
		assert.NoError(t, err)
		a, _ := fileentry.NewFileEntry("/media/video_tmp/20191102_test.json")
		b, _ := fileentry.NewFileEntry("/media/video_tmp/20191102_test.log")
		c, _ := fileentry.NewFileEntry("/media/video_tmp/20191102_test.mp4")
		expected := []fileentry.FileEntry{*a, *b, *c}
		assert.Equal(t, group.date, date)
		assert.EqualValues(t, group.entries, expected)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("invalid ext", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			date, _ := intdate.NewIntDate(20191102)

			group, err := NewEntryGroupFromTSPath(fs, date, "/media/video_tmp/20191102_test.mp4")
			assert.Error(t, err)
			assert.Nil(t, group)
		})
	})
}

func TestEntryGroup__EntryGroup__IsEmpty(t *testing.T) {
	date, _ := intdate.NewIntDate(20191102)

	t.Run("ok", func(t *testing.T) {
		t.Run("returns true", func(t *testing.T) {
			entries := []fileentry.FileEntry{}
			group, _ := NewEntryGroup(date, entries)

			actual := group.IsEmpty()
			assert.True(t, actual)
		})

		t.Run("returns false", func(t *testing.T) {
			entry, _ := fileentry.NewFileEntry("/media/video_tmp/20191102_test.mp4")
			entries := []fileentry.FileEntry{*entry}
			group, _ := NewEntryGroup(date, entries)

			actual := group.IsEmpty()
			assert.False(t, actual)
		})
	})
}
