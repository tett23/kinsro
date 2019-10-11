package writer

import (
	"testing"

	"github.com/tett23/kinsro/src/vindex/reader"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func TestAppend(t *testing.T) {
	item := vindexdata.NewBinaryIndex("foo")

	ch := writer.Append(item)
	<-ch

	ch = writer.Append(vindexdata.NewBinaryIndex("a"))
	<-ch

	ret, err := reader.ReadAll()
	if err != nil {
		t.Error(err)
	}

	t.Log(ret)
}
