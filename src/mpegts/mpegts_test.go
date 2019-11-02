package mpegts

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestMpegTS__NewMpegTS(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		actual, err := NewMpegTS("/media/video_tmp/20191102_test.ts")
		assert.NoError(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("invalid ext", func(t *testing.T) {
			actual, err := NewMpegTS("/media/video_tmp/20191102_test.mp4")
			assert.Error(t, err)
			assert.Nil(t, actual)
		})

		t.Run("invalid filename", func(t *testing.T) {
			actual, err := NewMpegTS("/media/video_tmp/test.ts")
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	})
}

func TestMpegTS__MpegTS__IsEncodable(t *testing.T) {
	tsPath := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, tsPath, []byte{}, 0744)
		ts, _ := NewMpegTS(tsPath)

		actual := ts.IsEncodable(fs)
		assert.True(t, actual)
	})

	t.Run("return false", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			ts, _ := NewMpegTS(tsPath)

			actual := ts.IsEncodable(fs)
			assert.False(t, actual)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, tsPath+".lock", []byte("2147483647"), 0744)
			ts, _ := NewMpegTS(tsPath)

			actual := ts.IsEncodable(fs)
			assert.False(t, actual)
		})
	})
}

func TestMpegTS__MpegTS__Dest(t *testing.T) {
	tsPath := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		ts, _ := NewMpegTS(tsPath)

		actual := ts.Dest()
		assert.Equal(t, actual, "/20191102_test.mp4")
	})
}

func TestMpegTS__MpegTS__Remove(t *testing.T) {
	tsPath := "/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		afero.WriteFile(fs, tsPath, []byte{}, 0744)
		ts, _ := NewMpegTS(tsPath)

		err := ts.Remove(fs)
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("file does not exists", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			ts, _ := NewMpegTS(tsPath)

			err := ts.Remove(fs)
			assert.Error(t, err)
		})

		t.Run("file locked", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			afero.WriteFile(fs, tsPath+".lock", []byte("2147483647"), 0744)
			ts, _ := NewMpegTS(tsPath)

			err := ts.Remove(fs)
			assert.Error(t, err)
		})
	})
}

func TestMpegTS__isMpegTSPath(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ok := isMpegTSPath("/media/video_tmp/20191102_test.ts")
		assert.True(t, ok)
	})

	t.Run("return false", func(t *testing.T) {
		t.Run("invalid ext", func(t *testing.T) {
			ok := isMpegTSPath("/media/video_tmp/20191102_test.mp4")
			assert.False(t, ok)
		})

		t.Run("invalid filename", func(t *testing.T) {
			ok := isMpegTSPath("/media/video_tmp/test.ts")
			assert.False(t, ok)
		})
	})
}
