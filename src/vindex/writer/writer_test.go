package writer

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestAppend(t *testing.T) {
	conf := config.Config{
		VIndexPath: "/test/vindex",
	}
	item1 := vindexdata.NewVIndexItem("video1", 20190101, "test1.ts")
	ch := Append(&conf, &item1)
	<-ch

	item2 := vindexdata.NewVIndexItem("video1", 20190101, "test2.ts")
	ch = Append(&conf, &item2)
	<-ch

	ret, err := reader.ReadAll(conf.VIndexPath)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(vindexdata.VIndex{item1, item2}, ret); diff != "" {
		t.Error(diff)
	}
}
