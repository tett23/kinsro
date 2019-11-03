package moveentries

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/entrygroup"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/intdate"
	"github.com/tett23/kinsro/src/storages"
)

func TestMoveEntryGroup(t *testing.T) {
	date, _ := intdate.NewIntDate(20201010)

	t.Run("ok", func(t *testing.T) {
		t.Run("copy file", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			content := "test"
			src, _ := fileentry.NewFileEntry("/src.mp4")
			afero.WriteFile(fs, src.Src(), []byte(content), 0744)
			items := []fileentry.FileEntry{*src}
			group, _ := entrygroup.NewEntryGroup(date, items)
			storage := storages.NewStorage("/media/video1")

			err := MoveEntryGroup(fs, group, storage)
			assert.NoError(t, err)

			ok, _ := afero.Exists(fs, src.Src())
			assert.False(t, ok)

			ok, _ = afero.Exists(fs, "/media/video1/2020/10/10/src.mp4")
			assert.True(t, ok)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("src does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			src, _ := fileentry.NewFileEntry("/src.mp4")
			items := []fileentry.FileEntry{*src}
			group, _ := entrygroup.NewEntryGroup(date, items)
			storage := storages.NewStorage("/media/video1")

			err := MoveEntryGroup(fs, group, storage)
			assert.Error(t, err)
		})
	})
}

func Test__MoveEntries__Copy(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("copy file", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			content := "test"
			src, _ := fileentry.NewFileEntry("/src.mp4")
			dest, _ := fileentry.NewFileEntry("/dest.mp4")
			afero.WriteFile(fs, src.Src(), []byte(content), 0744)

			err := Copy(fs, src, dest)
			assert.NoError(t, err)

			actual, _ := afero.ReadFile(fs, dest.Src())
			assert.Equal(t, string(actual), content)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("src path does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			src, _ := fileentry.NewFileEntry("/src.mp4")
			dest, _ := fileentry.NewFileEntry("/dest.mp4")

			err := Copy(fs, src, dest)
			assert.Error(t, err)
		})
	})
}
