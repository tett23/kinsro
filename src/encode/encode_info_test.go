package encode_test

import (
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/encode"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestEncode__NewEncodeInfo(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		conf := config.Config{
			StoragePaths: []string{"/out"},
		}
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()
		fs.MkdirAll("/test", 0744)
		fs.MkdirAll("/out", 0744)

		info, err := encode.NewEncodeInfo(&conf, fs, "/test/201910240105_test.ts")
		assert.Nil(t, err)
		assert.NotNil(t, info)
	})

	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		conf := config.Config{
			StoragePaths: []string{"/out"},
		}
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()

		info, err := encode.NewEncodeInfo(&conf, fs, "/test/test.ts")
		assert.NotNil(t, err)
		assert.Nil(t, info)
	})
}

func TestEncode__EncodeInfo__Prefix(t *testing.T) {
	filesystem.ResetTestFs()
	t.Run("", func(t *testing.T) {
		info := encode.EncodeInfo{
			RawPath: "/foo/201910240105_test.ts",
		}
		assert.Equal(t, info.Prefix(), "201910240105_test")
	})
}

func TestEncode__EncodeInfo__TSDest(t *testing.T) {
	filesystem.ResetTestFs()
	t.Run("", func(t *testing.T) {
		info := encode.EncodeInfo{
			RawPath: "/foo/201910240105_test.ts",
		}
		assert.Equal(t, info.TSDest(), "/foo/201910240105_test.mp4")
	})
}

func TestEncode__EncodeInfo__MoveFiles(t *testing.T) {
	filesystem.ResetTestFs()
	t.Run("", func(t *testing.T) {
		info := encode.EncodeInfo{
			RawPath: "/foo/201910240105_test.ts",
		}

		actual := info.MoveFiles()
		expected := []string{"201910240105_test.mp4"}
		assert.Equal(t, actual, expected)
	})

	t.Run("", func(t *testing.T) {
		info := encode.EncodeInfo{
			RawPath:           "/foo/201910240105_test.ts",
			MetadataFileNames: []string{"201910240105_test.mp4"},
		}

		actual := info.MoveFiles()
		expected := []string{"201910240105_test.mp4"}
		assert.Equal(t, actual, expected)
	})
}

func TestEncode__EncodeInfo__Move(t *testing.T) {
	statfs := func(path string, stat *syscall.Statfs_t) error {
		stat.Bsize = 4086
		stat.Bfree = 1000

		return nil
	}

	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()
		fs.MkdirAll("/foo", 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.ts", []byte{}, 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.json", []byte{}, 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.mp4", []byte{}, 0744)

		info := encode.EncodeInfo{
			RawPath:           "/foo/201910240105_test.ts",
			MetadataFileNames: []string{"201910240105_test.json"},
			StoragePaths:      []string{"/out"},
		}
		err := info.Move(statfs, fs)
		assert.Nil(t, err)

		files, err := afero.ReadDir(fs, "/out/2019/10/24")
		moved := []string{}
		for i := range files {
			moved = append(moved, files[i].Name())
		}

		expected := []string{
			"201910240105_test.json",
			"201910240105_test.mp4",
		}
		assert.Equal(t, moved, expected)
	})
}

func TestEncode__EncodeInfo__FetchMeta(t *testing.T) {
	t.Run("", func(t *testing.T) {
		filesystem.ResetTestFs()
		fs := filesystem.GetFs()
		fs.MkdirAll("/foo", 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.ts", []byte{}, 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.json", []byte{}, 0744)
		afero.WriteFile(fs, "/foo/201910240105_test.ts.program.txt", []byte{}, 0744)
		fs.MkdirAll("/foo/201910240105_test.dir", 0744)
		afero.WriteFile(fs, "/foo/202010240105_test.ts.program.txt", []byte{}, 0744)

		info := encode.EncodeInfo{
			RawPath: "/foo/201910240105_test.ts",
		}

		actual, err := info.FetchMeta(fs)
		assert.Nil(t, err)
		expected := []string{"201910240105_test.json", "201910240105_test.ts.program.txt"}
		assert.Equal(t, actual, expected)
	})
}
