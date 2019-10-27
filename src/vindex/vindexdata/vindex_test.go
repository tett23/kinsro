package vindexdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
)

func TestMain(m *testing.M) {
	filesystem.CopyDotfile()
	config.Initialize()
}

func TestVIndexToBinaryAndNewIndexFromBinary(t *testing.T) {
	vindex := VIndex{
		NewVIndexItem("video1", uint64(20200101), "test1.ts"),
		NewVIndexItem("video1", uint64(20200102), "test2.ts"),
	}

	bin, err := vindex.ToBinary()
	if err != nil {
		t.Error(err)
	}

	newVindex, err := NewVIndexFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(vindex, newVindex); diff != "" {
		t.Errorf(diff)
	}
}
