package metafile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetaFile__NewMetaFile(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		actual := NewMetaFile("/media/video_tmp/20191102_test.ts")

		assert.NotNil(t, actual)
	})
}
