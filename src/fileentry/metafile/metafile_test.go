package metafile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetaFile__NewMetaFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		actual, err := NewMetaFile("/media/video_tmp/20191102_test.json")

		assert.NoError(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("error", func(t *testing.T) {
			actual, err := NewMetaFile("/test.ts")

			assert.Error(t, err)
			assert.Nil(t, actual)
		})

		t.Run("error", func(t *testing.T) {
			actual, err := NewMetaFile("")

			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	})
}

func TestMetaFile__isValidMetaFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		ok := isValidMetaFile("/media/video_tmp/20191102_test.json")
		assert.True(t, ok)
	})

	t.Run("error", func(t *testing.T) {
		t.Run("ts", func(t *testing.T) {
			ok := isValidMetaFile("/test.ts")
			assert.False(t, ok)
		})

		t.Run("mp4", func(t *testing.T) {
			ok := isValidMetaFile("/test.mp4")
			assert.False(t, ok)
		})
	})
}
