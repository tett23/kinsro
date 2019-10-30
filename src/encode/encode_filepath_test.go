package encode_test

import (
	"syscall"
	"testing"

	"github.com/tett23/kinsro/src/encode"
	"gotest.tools/assert"
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
		assert.NilError(t, err)
		assert.Equal(t, file.Storage, "/fuga")
	})

	t.Run("", func(t *testing.T) {
		statfs := func(path string, stat *syscall.Statfs_t) error {
			stat.Bsize = 4086
			stat.Bfree = 1000

			return nil
		}

		file, err := encode.NewEncodeFilePath(statfs, "/test", []string{"/hoge"})
		assert.NilError(t, err)
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
