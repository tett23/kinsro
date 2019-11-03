package commands

import (
	"syscall"
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/syscalls"
	"github.com/tett23/kinsro/src/tests"
)

func buildStatfs() syscalls.Statfs {
	count := 0

	return func(path string, stat *syscall.Statfs_t) error {
		stat.Bsize = 4086
		stat.Bfree = uint64(1000 * (count + 1))
		count++

		return nil
	}
}

func TestCommands__EncodeTS(t *testing.T) {
	tests.Setup()
	ss := syscalls.Syscalls{
		Statfs:  buildStatfs(),
		Chdir:   func(_ string) error { return nil },
		Symlink: func(_, _ string) error { return nil },
	}

	t.Run("ok", func(t *testing.T) {
		t.Run("copy file", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			conf := config.GetConfig()

			path := "/media/video_tmp/20191102_test.ts"
			fs.MkdirAll("/media/video_tmp", 0755)
			afero.WriteFile(fs, path, []byte{}, 0744)
			options := EncodeOptions{
				RemoveTS: true,
			}

			err := EncodeTS(conf, &ss, fs, path, options)
			assert.NoError(t, err)

			ok, _ := afero.Exists(fs, path)
			assert.False(t, ok)
		})

		t.Run("copy file", func(t *testing.T) {
			fs := filesystem.ResetTestFs()

			conf := config.GetConfig()

			path := "/media/video_tmp/20191102_test.ts"
			fs.MkdirAll("/media/video_tmp", 0755)
			afero.WriteFile(fs, path, []byte{}, 0744)
			options := EncodeOptions{
				RemoveTS: false,
			}

			err := EncodeTS(conf, &ss, fs, path, options)
			assert.NoError(t, err)

			ok, _ := afero.Exists(fs, path)
			assert.True(t, ok)
		})
	})
}
