package vindexdata_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestVIndexToBinaryAndNewIndexFromBinary(t *testing.T) {
	vindex := vindexdata.VIndex{
		vindexdata.NewVIndexItem("video1", uint64(20200101), "test1.ts"),
		vindexdata.NewVIndexItem("video1", uint64(20200102), "test2.ts"),
	}

	bin, err := vindex.ToBinary()
	if err != nil {
		t.Error(err)
	}

	newVindex, err := vindexdata.NewVIndexFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(vindex, newVindex); diff != "" {
		t.Errorf(diff)
	}
}
