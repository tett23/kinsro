package vindexdata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVIndexData__ParseFilepath(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		path := "/test/2019/10/10/test.mp4"
		storagePaths := []string{"/test"}
		actual, err := ParseFilepath(storagePaths, path)
		assert.NoError(t, err)

		expected, err := NewVIndexItem("test", 20191010, "test.mp4")
		assert.NoError(t, err)
		assert.Equal(t, actual, expected)
	})

	t.Run("when passed invalid ext", func(t *testing.T) {
		path := "/test/2019/10/10/test.ts"
		storagePaths := []string{"/test"}

		actual, err := ParseFilepath(storagePaths, path)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("returns error when storage not found", func(t *testing.T) {
		path := "/test/2019/10/10/test.mp4"
		storagePaths := []string{""}

		actual, err := ParseFilepath(storagePaths, path)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("returns error when storage not found", func(t *testing.T) {
		path := "/test/2019/10/10/test.ts"
		storagePaths := []string{"/hoge"}

		actual, err := ParseFilepath(storagePaths, path)
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}
