package vindexdata

import (
	"reflect"
	"testing"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestBinaryIndexItemToBinary(t *testing.T) {
	item := vindexdata.NewBinaryIndexItem("foo", 20200101)
	bin := item.ToBinary()
	restore, err := vindexdata.NewBinaryIndexItemFromBinary(bin)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(item, restore) {
		t.Errorf("")
	}
}
