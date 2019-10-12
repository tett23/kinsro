package vindexdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVIndexItemToBinaryAndNewVIndexItemFromBinary(t *testing.T) {
	item := NewVIndexItem("video1", uint64(20200101), "test.ts")
	bin := item.ToBinary()
	restore, err := NewBinaryIndexItemFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(item, restore); diff != "" {
		t.Errorf(diff)
	}
}
