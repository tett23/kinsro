package vindexdata_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestVIndexToBinaryAndNewIndexFromBinary(t *testing.T) {
	item1, err := vindexdata.NewVIndexItem("video1", 20200101, "test1.ts")
	assert.NoError(t, err)

	item2, err := vindexdata.NewVIndexItem("video1", 20200102, "test2.ts")
	assert.NoError(t, err)

	vindex := vindexdata.VIndex{*item1, *item2}

	bin, err := vindex.ToBinary()
	assert.NoError(t, err)

	newVindex, err := vindexdata.NewVIndexFromBinary(bin)
	assert.NoError(t, err)
	assert.Equal(t, vindex, newVindex)
}
