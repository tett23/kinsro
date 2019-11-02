package writer

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func TestAppend(t *testing.T) {
	conf := config.Config{
		VIndexPath: "/test/vindex",
	}
	item1, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.ts")
	assert.NoError(t, err)
	ch := Append(&conf, item1)
	<-ch

	item2, err := vindexdata.NewVIndexItem("video1", 20190101, "test2.ts")
	assert.NoError(t, err)
	ch = Append(&conf, item2)
	<-ch

	ret, err := reader.ReadAll(conf.VIndexPath)
	assert.NoError(t, err)
	assert.Equal(t, vindexdata.VIndex{*item1, *item2}, ret)
}
