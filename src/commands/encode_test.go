package commands

import (
	"syscall"
	"testing"
	"time"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/clock"
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
		t.Run("removes TS file", func(t *testing.T) {
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

		t.Run("does not removes TS file", func(t *testing.T) {
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

func TestCommands__EncodeTSAll(t *testing.T) {
	tests.Setup()
	ss := syscalls.Syscalls{
		Statfs:  buildStatfs(),
		Chdir:   func(_ string) error { return nil },
		Symlink: func(_, _ string) error { return nil },
	}
	videoTmpPath := "/media/video_tmp"

	t.Run("ok", func(t *testing.T) {
		t.Run("returns no error", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			conf := config.GetConfig()
			fs.MkdirAll(videoTmpPath, 0755)
			ts1 := videoTmpPath + "/20191102_test.ts"
			afero.WriteFile(fs, ts1, []byte{}, 0744)
			ts2 := videoTmpPath + "/20191103_test.ts"
			afero.WriteFile(fs, ts2, []byte{}, 0744)
			options := EncodeOptions{RemoveTS: true}

			err := EncodeTSAll(conf, &ss, fs, videoTmpPath, options)
			assert.NoError(t, err)
		})
	})
}

func TestCommands__headTS(t *testing.T) {
	videoTmpPath := "/media/video_tmp"

	t.Run("ok", func(t *testing.T) {
		t.Run("returns TS path", func(t *testing.T) {
			clock.Set(time.Now())
			defer clock.Reset()
			fs := filesystem.ResetTestFs()
			fs.MkdirAll(videoTmpPath, 0755)
			ts1 := videoTmpPath + "/20191102_test.ts"
			afero.WriteFile(fs, ts1, []byte{}, 0744)
			ts2 := videoTmpPath + "/20191103_test.ts"
			afero.WriteFile(fs, ts2, []byte{}, 0744)

			ts, ok, err := headTS(fs, videoTmpPath, []string{})
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, ts, ts1)
		})

		t.Run("filter by filename", func(t *testing.T) {
			clock.Set(time.Now())
			defer clock.Reset()
			fs := filesystem.ResetTestFs()
			fs.MkdirAll(videoTmpPath, 0755)
			ts1 := videoTmpPath + "/test1.ts"
			afero.WriteFile(fs, ts1, []byte{}, 0744)
			ts2 := videoTmpPath + "/test2.ts"
			afero.WriteFile(fs, ts2, []byte{}, 0744)

			ts, ok, err := headTS(fs, videoTmpPath, []string{})
			assert.NoError(t, err)
			assert.False(t, ok)
			assert.Equal(t, ts, "")
		})

		t.Run("filter by ignore paths", func(t *testing.T) {
			clock.Set(time.Now())
			defer clock.Reset()
			fs := filesystem.ResetTestFs()
			fs.MkdirAll(videoTmpPath, 0755)
			ts1 := videoTmpPath + "/20191102_test.ts"
			afero.WriteFile(fs, ts1, []byte{}, 0744)
			ts2 := videoTmpPath + "/20191103_test.ts"
			afero.WriteFile(fs, ts2, []byte{}, 0744)

			ts, ok, err := headTS(fs, videoTmpPath, []string{ts1})
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, ts, ts2)
		})

		t.Run("filter by current time", func(t *testing.T) {
			clock.Set(time.Date(2000, 1, 10, 0, 0, 0, 0, time.Local))
			defer clock.Reset()

			fs := filesystem.ResetTestFs()
			fs.MkdirAll(videoTmpPath, 0755)
			ts1 := videoTmpPath + "/20000110_test1.ts"
			afero.WriteFile(fs, ts1, []byte{}, 0744)
			ts2 := videoTmpPath + "/20000107_test1.ts"
			afero.WriteFile(fs, ts2, []byte{}, 0744)

			ts, ok, err := headTS(fs, videoTmpPath, []string{})
			assert.NoError(t, err)
			assert.True(t, ok)
			assert.Equal(t, ts, ts2)
		})
	})
}
