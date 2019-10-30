package encode_test

import (
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestEncode__NewEncodeFilePath(t *testing.T) {
	t.Run("", func(t *testing.T) {
		count := 0
		statfs := func(path string, stat *syscall.Statfs_t) error {
			stat.Bsize = 4086
			stat.Bfree = uint64(1000 * (count + 1))
			count++

			return nil
		}

		file, err := encode.NewEncodeFilePath(statfs, "/test", []string{"/hoge", "/fuga"})
		assert.NoError(t, err)
		assert.Equal(t, file.Storage, "/fuga")
	})

	t.Run("", func(t *testing.T) {
		statfs := func(path string, stat *syscall.Statfs_t) error {
			stat.Bsize = 4086
			stat.Bfree = 1000

			return nil
		}

		file, err := encode.NewEncodeFilePath(statfs, "/test", []string{"/hoge"})
		assert.NoError(t, err)
		assert.Equal(t, file.Storage, "/hoge")
	})
}

func TestEncode__EncodeFilePath__Dest(t *testing.T) {
	item := encode.EncodeFilePath{
		Path:    "/foo/201910240105_test.mp4",
		Storage: "/bar",
	}

	assert.Equal(t, item.Dest(), "/bar/2019/10/24/201910240105_test.mp4")
}

func TestEncode__EncodeFilePath__Dir(t *testing.T) {
	item := encode.EncodeFilePath{
		Path:    "/foo/201910240105_test.mp4",
		Storage: "/bar",
	}

	assert.Equal(t, item.Dir(), "/bar/2019/10/24")
}

func TestEncode__EncodeFilePath__Copy(t *testing.T) {
	fs := filesystem.ResetTestFs()
	content := "test"
	afero.WriteFile(fs, "/test/201910240105_test.txt", []byte(content), 0744)
	item := encode.EncodeFilePath{
		Path:    "/test/201910240105_test.txt",
		Storage: "/bar",
	}

	err := item.Copy(fs)
	assert.NoError(t, err)

	ok, _ := afero.Exists(fs, item.Src())
	assert.True(t, ok)

	ok, _ = afero.Exists(fs, item.Dest())
	assert.True(t, ok)

	actual, _ := afero.ReadFile(fs, item.Dest())
	assert.Equal(t, string(actual), content)
}

func TestEncode__EncodeFilePath__ToVIndexItem(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		item := encode.EncodeFilePath{
			Path:    "/test/201910240105_test.mp4",
			Storage: "/bar",
		}

		vindexItem, err := item.ToVIndexItem()
		assert.NoError(t, err)
		assert.NotNil(t, vindexItem)
	})

	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		item := encode.EncodeFilePath{
			Path:    "/test/201910240105_test.txt",
			Storage: "/bar",
		}

		vindexItem, err := item.ToVIndexItem()
		assert.Error(t, err)
		assert.Nil(t, vindexItem)
	})
}
