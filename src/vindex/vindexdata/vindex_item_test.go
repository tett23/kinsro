package vindexdata_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestVIndex__NewVIndexItemFromBinary(t *testing.T) {
	item := vindexdata.NewVIndexItem("video1", uint64(20200101), "test.mp4")
	bin := item.ToBinary()
	restore, err := vindexdata.NewBinaryIndexItemFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(item, *restore); diff != "" {
		t.Errorf(diff)
	}
}

func TestVIndex__VIndexItem__FullPath(t *testing.T) {
	t.Run("strgage exists", func(t *testing.T) {
		item := vindexdata.NewVIndexItem("video1", uint64(20201231), "test.mp4")

		actual, err := item.FullPath([]string{"/media/video1"})
		assert.NoError(t, err)
		assert.Equal(t, actual, "/media/video1/2020/12/31/test.mp4")
	})

	t.Run("storage does not exists", func(t *testing.T) {
		item := vindexdata.NewVIndexItem("does_not_exists", uint64(20201231), "test.mp4")

		actual, err := item.FullPath([]string{"/media/video1"})
		assert.Error(t, err)
		assert.Equal(t, actual, "")
	})
}

func TestVIndex__VIndexItem__Path(t *testing.T) {
	item := vindexdata.NewVIndexItem("video1", uint64(20201231), "test.mp4")

	actual := item.Path()
	assert.Equal(t, actual, "2020/12/31/test.mp4")
}
