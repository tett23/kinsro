package vindexdata_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestVIndexItemToBinaryAndNewVIndexItemFromBinary(t *testing.T) {
	item := vindexdata.NewVIndexItem("video1", uint64(20200101), "test.ts")
	bin := item.ToBinary()
	restore, err := vindexdata.NewBinaryIndexItemFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(item, *restore); diff != "" {
		t.Errorf(diff)
	}
}
