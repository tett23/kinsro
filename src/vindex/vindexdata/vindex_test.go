package vindexdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestVIndexToBinaryAndNewIndexFromBinary(t *testing.T) {
	vindex := VIndex{
		NewVIndexItem("foo", uint64(20200101)),
		NewVIndexItem("foo", uint64(20200101)),
	}
	//
	bin, err := vindex.ToBinary()
	if err != nil {
		t.Error(err)
	}

	newVindex, err := NewVIndexFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(vindex, newVindex); diff != "" {
		t.Errorf((diff))
	}
}
