package mpegts

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/fileentry/metafile"
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

		t.Run("error", func(t *testing.T) {
			actual, err := NewMpegTS("")

			assert.Error(t, err)
			assert.Nil(t, actual)
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

func TestMpegTS__MpegTS__MetaFiles(t *testing.T) {
	path := "/media/video_tmp/20191102_test.ts"
	createFiles := func(fs afero.Fs, files []string) {
		for i := range files {
			afero.WriteFile(fs, files[i], []byte{}, 0644)
		}
	}

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		ts, _ := NewMpegTS(path)
		createFiles(fs, []string{
			path,
			"/media/video_tmp/20191102_test.ts.lock",
			"/media/video_tmp/20191102_test.mp4",
			"/media/video_tmp/20191102_test.json",
			"/media/video_tmp/20191102_test.log",
			"/media/video_tmp/20191103_test.log",
		})

		actual, err := ts.MetaFiles(fs)
		assert.NoError(t, err)
		a, _ := metafile.NewMetaFile("/media/video_tmp/20191102_test.json")
		b, _ := metafile.NewMetaFile("/media/video_tmp/20191102_test.log")
		assert.Equal(t, actual, []metafile.MetaFile{*a, *b})
	})

	t.Run("return false", func(t *testing.T) {
		t.Run("invalid ext", func(t *testing.T) {
		})

		t.Run("invalid filename", func(t *testing.T) {
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
