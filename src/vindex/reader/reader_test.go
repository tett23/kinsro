package reader_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

var conf = config.Config{
	VIndexPath: "/test/vindex",
}

func TestFindByFilename1(t *testing.T) {
	item, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
	assert.NoError(t, err)

	vindex := vindexdata.VIndex{*item}
	err = writer.CreateNewIndexFile(&conf, vindex)
	assert.NoError(t, err)

	record, err := reader.FindByFilename(&conf, vindex[0].Filename)
	assert.NoError(t, err)
	assert.Equal(t, record, item)
}

func TestFindByFilename2(t *testing.T) {
	item, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
	assert.NoError(t, err)

	vindex := vindexdata.VIndex{*item}
	err = writer.CreateNewIndexFile(&conf, vindex)
	assert.NoError(t, err)

	record, err := reader.FindByFilename(&conf, "test2.mp4")
	assert.NoError(t, err)
	assert.Nil(t, record)
}

func TestFindByDigest1(t *testing.T) {
	item, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
	assert.NoError(t, err)

	vindex := vindexdata.VIndex{*item}
	err = writer.CreateNewIndexFile(&conf, vindex)
	assert.NoError(t, err)

	record, err := reader.FindByDigest(&conf, vindex[0].Digest.Hex())
	assert.NoError(t, err)
	assert.Equal(t, record, item)
}

func TestFindByDigest2(t *testing.T) {
	item, err := vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4")
	assert.NoError(t, err)

	vindex := vindexdata.VIndex{*item}
	err = writer.CreateNewIndexFile(&conf, vindex)
	assert.NoError(t, err)

	record, err := reader.FindByDigest(&conf, "a4ddbc2cc9364c45b0012d9acdf32cf0")
	assert.NoError(t, err)
	assert.Nil(t, record)
}
