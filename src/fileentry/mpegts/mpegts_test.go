package mpegts

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/fileentry"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/intdate"
	"github.com/tett23/kinsro/src/storages"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
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

func TestMpegTS__MpegTS__ToEntryGroup(t *testing.T) {
	path := "/media/video_tmp/20191102_test.ts"

	t.Run("ok", func(t *testing.T) {
		fs := filesystem.ResetTestFs()
		ts, _ := NewMpegTS(path)

		actual, err := ts.ToEntryGroup(fs)
		assert.NoError(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("invalid date", func(t *testing.T) {
			fs := filesystem.ResetTestFs()
			entry, _ := fileentry.NewFileEntry("/media/video_tmp/hogehoge.ts")
			ts := MpegTS{
				FileEntry: entry,
			}

			actual, err := ts.ToEntryGroup(fs)
			assert.Error(t, err)
			assert.Nil(t, actual)
		})

	})
}

func TestMpegTS__MpegTS__ToIntDate(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ts, _ := NewMpegTS("/media/video_tmp/20191102_test.ts")

		actual, err := ts.ToIntDate()
		assert.NoError(t, err)
		assert.Equal(t, actual, intdate.IntDate(20191102))
	})

	t.Run("error", func(t *testing.T) {
		t.Run("invalid filename", func(t *testing.T) {
			entry, _ := fileentry.NewFileEntry("/media/video_tmp/test.ts")
			ts := MpegTS{
				FileEntry: entry,
			}

			_, err := ts.ToIntDate()
			assert.Error(t, err)
		})

		t.Run("invalid filename", func(t *testing.T) {
			entry, _ := fileentry.NewFileEntry("/media/video_tmp/hogehoge.ts")
			ts := MpegTS{
				FileEntry: entry,
			}

			actual, err := ts.ToIntDate()
			assert.Error(t, err)
			assert.Equal(t, actual, intdate.IntDate(-1))
		})
	})
}

func TestMpegTS__MpegTS__ToVIndexItem(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ts, _ := NewMpegTS("/media/video_tmp/20191102_test.ts")
		storage := storages.NewStorage("/media/video1")

		actual, err := ts.ToVIndexItem(storage)
		assert.NoError(t, err)
		assert.NotNil(t, actual)
		expected, _ := vindexdata.NewVIndexItem("video1", 20191102, "20191102_test.mp4")
		assert.Equal(t, actual, expected)
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
