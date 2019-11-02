package moventries

import (
	"errors"
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/fileentry/entrygroup"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/intdate"
)

func buildStatfs() Statfs {
	count := 0

	return func(path string, stat *syscall.Statfs_t) error {
		stat.Bsize = 4086
		stat.Bfree = uint64(1000 * (count + 1))
		count++

		return nil
	}
}

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

			err := MoveEntryGroup(fs, buildStatfs(), group, []string{"/storage"})
			assert.NoError(t, err)

			ok, _ := afero.Exists(fs, src.Src())
			assert.False(t, ok)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("no storages", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			src, _ := fileentry.NewFileEntry("/src.mp4")
			items := []fileentry.FileEntry{*src}
			group, _ := entrygroup.NewEntryGroup(date, items)

			err := MoveEntryGroup(fs, buildStatfs(), group, []string{})
			assert.Error(t, err)
		})

		t.Run("src does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			src, _ := fileentry.NewFileEntry("/src.mp4")
			items := []fileentry.FileEntry{*src}
			group, _ := entrygroup.NewEntryGroup(date, items)

			err := MoveEntryGroup(fs, buildStatfs(), group, []string{"/storage"})
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

func TestMoveEntries__mostSpacefulStorage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage path", func(t *testing.T) {
			storage, err := mostSpacefulStorage(buildStatfs(), []string{"/storage1"})
			assert.NoError(t, err)
			assert.Equal(t, storage, "/storage1")
		})

		t.Run("return most spaceful storage", func(t *testing.T) {
			storage, err := mostSpacefulStorage(buildStatfs(), []string{"/storage1", "/storage2"})
			assert.NoError(t, err)
			assert.Equal(t, storage, "/storage2")
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("no storages", func(t *testing.T) {
			storage, err := mostSpacefulStorage(buildStatfs(), []string{})
			assert.Error(t, err)
			assert.Equal(t, storage, "")
		})

		t.Run("statfs returns error", func(t *testing.T) {
			statfs := func(path string, stat *syscall.Statfs_t) error {
				return errors.New("")
			}

			storage, err := mostSpacefulStorage(statfs, []string{"/storage1", "/storage2"})
			assert.Error(t, err)
			assert.Equal(t, storage, "")
		})
	})
}
