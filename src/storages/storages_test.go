package storages

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorages__Storages__FindByName(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage", func(t *testing.T) {
			items := NewStoragesFromPaths([]string{"/media/video1", "/media/video2"})

			actual, ok := items.FindByName("video1")
			assert.True(t, ok)
			assert.Equal(t, actual.Path, "/media/video1")
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("no storages", func(t *testing.T) {
			items := NewStoragesFromPaths([]string{"/media/video1", "/media/video2"})

			actual, ok := items.FindByName("does_not_exists")
			assert.False(t, ok)
			assert.Nil(t, actual)
		})
	})
}

func TestStorages__Storages__MostSpacefulStorage(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		t.Run("returns storage", func(t *testing.T) {
			items := NewStoragesFromPaths([]string{"/media/video1", "/media/video2"})

			actual, err := items.MostSpacefulStorage(buildStatfs())
			assert.NoError(t, err)
			assert.Equal(t, actual.Path, "/media/video2")
		})
	})

	t.Run("error", func(t *testing.T) {
		t.Run("no storages", func(t *testing.T) {
			items := NewStoragesFromPaths([]string{})

			actual, err := items.MostSpacefulStorage(buildStatfs())
			assert.Error(t, err)
			assert.Nil(t, actual)
		})
	})
}
