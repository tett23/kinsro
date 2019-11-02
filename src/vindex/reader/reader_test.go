package reader

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

var conf = config.Config{
	VIndexPath:   "/test/vindex",
	StoragePaths: []string{"/media/video1"},
}

func TestReader__ReadAll(t *testing.T) {
	conf := config.Config{
		VIndexPath:   "/test/vindex",
		StoragePaths: []string{"/media/video1"},
	}

	t.Run("ok", func(t *testing.T) {
		t.Run("returns VIndex", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			item, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
			assert.NoError(t, err)

			vindex := vindexdata.VIndex{*item}
			writer.CreateNewIndexFile(&conf, vindex)

			actual, err := ReadAll(fs, conf.VIndexPath)
			assert.NoError(t, err)
			assert.Equal(t, actual, vindex)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("vindex does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			actual, err := ReadAll(fs, conf.VIndexPath)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	})
}

func TestReader__FindByFullPath(t *testing.T) {
	conf := config.Config{
		VIndexPath:   "/test/vindex",
		StoragePaths: []string{"/media/video1"},
	}

	t.Run("ok", func(t *testing.T) {
		t.Run("returns VIndexItem", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
			vindex := vindexdata.VIndex{*item}
			writer.CreateNewIndexFile(&conf, vindex)
			fullPath, _ := item.FullPath(conf.StoragePaths)

			record, ok, err := FindByFullPath(&conf, fs, fullPath)
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, record, item)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("path error", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")

			vindex := vindexdata.VIndex{*item}
			writer.CreateNewIndexFile(&conf, vindex)

			record, ok, err := FindByFullPath(&conf, fs, "/test2.mp4")
			assert.Error(t, err)
			assert.False(t, ok)
			assert.Nil(t, record)
		})
	})
}

func TestReader__FindByDigest(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns VindexItem", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
			vindex := vindexdata.VIndex{*item}
			writer.CreateNewIndexFile(&conf, vindex)

			record, ok, err := FindByDigest(&conf, fs, item.HexDigest())
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, record, item)
		})

		t.Run("no vindex", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")

			record, ok, err := FindByDigest(&conf, fs, item.HexDigest())
			assert.NoError(t, err)
			assert.False(t, ok)
			assert.Nil(t, record)
		})

		t.Run("not found", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
			vindex := vindexdata.VIndex{*item}
			writer.CreateNewIndexFile(&conf, vindex)

			record, ok, err := FindByDigest(&conf, fs, "does_not_exists")
			assert.NoError(t, err)
			assert.False(t, ok)
			assert.Nil(t, record)
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("invalid vindex", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, conf.VIndexPath, []byte("hoge"), 0644)
			item, _ := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")

			record, ok, err := FindByDigest(&conf, fs, item.HexDigest())
			assert.Error(t, err)
			assert.False(t, ok)
			assert.Nil(t, record)
		})
	})
}
