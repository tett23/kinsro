package vindexdata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestVIndex__NewVIndexItem(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		actual, err := vindexdata.NewVIndexItem("video1", 20200101, "test.mp4")
		assert.NoError(t, err)
		assert.NotNil(t, actual)
	})

	t.Run("check storage", func(t *testing.T) {
		actual, err := vindexdata.NewVIndexItem("/media/video1", 20200101, "test.mp4")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("check filename", func(t *testing.T) {
		actual, err := vindexdata.NewVIndexItem("video1", 20200101, "/media/video1/2020/01/01/test.mp4")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})

	t.Run("check date", func(t *testing.T) {
		actual, err := vindexdata.NewVIndexItem("video1", 1, "test.mp4")
		assert.Error(t, err)
		assert.Nil(t, actual)
	})
}

func TestVIndex__NewVIndexItemFromBinary(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		item, _ := vindexdata.NewVIndexItem("video1", 20200101, "test.mp4")
		bin := item.ToBinary()

		restore, err := vindexdata.NewBinaryIndexItemFromBinary(bin)
		assert.NoError(t, err)
		assert.Equal(t, item, restore)
	})

	t.Run("error", func(t *testing.T) {
		restore, err := vindexdata.NewBinaryIndexItemFromBinary([]byte{})
		assert.Error(t, err)
		assert.Nil(t, restore)
	})
}

func TestVIndex__VIndexItem__FullPath(t *testing.T) {
	t.Run("strgage exists", func(t *testing.T) {
		item, _ := vindexdata.NewVIndexItem("video1", 20201231, "test.mp4")

		actual, err := item.FullPath([]string{"/media/video1"})
		assert.NoError(t, err)
		assert.Equal(t, actual, "/media/video1/2020/12/31/test.mp4")
	})

	t.Run("storage does not exists", func(t *testing.T) {
		item, _ := vindexdata.NewVIndexItem("does_not_exists", 20201231, "test.mp4")

		actual, err := item.FullPath([]string{"/media/video1"})
		assert.Error(t, err)
		assert.Equal(t, actual, "")
	})
}

func TestVIndex__VIndexItem__Path(t *testing.T) {
	t.Run("returns path string", func(t *testing.T) {
		item, _ := vindexdata.NewVIndexItem("video1", 20201231, "test.mp4")

		actual := item.Path()
		assert.Equal(t, actual, "2020/12/31/test.mp4")
	})

	t.Run("pad with zero", func(t *testing.T) {
		item, _ := vindexdata.NewVIndexItem("video1", 20200101, "test.mp4")

		actual := item.Path()
		assert.Equal(t, actual, "2020/01/01/test.mp4")
	})
}
