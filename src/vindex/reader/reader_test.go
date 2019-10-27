package reader_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
	"github.com/tett23/kinsro/src/vindex/reader"
	"github.com/tett23/kinsro/src/vindex/vindexdata"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func TestMain(m *testing.M) {
	filesystem.CopyDotfile()
	config.Initialize()
}

var conf = config.Config{
	VIndexPath: "/test/vindex",
}

func TestFindByFilename1(t *testing.T) {
	vindex := vindexdata.VIndex{
		vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4"),
	}
	err := writer.CreateNewIndexFile(&conf, vindex)
	if err != nil {
		t.Error(err)
	}

	record, err := reader.FindByFilename(&conf, vindex[0].Filename)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(*record, vindex[0]); diff != "" {
		t.Error(diff)
	}
}

func TestFindByFilename2(t *testing.T) {
	vindex := vindexdata.VIndex{
		vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4"),
	}
	err := writer.CreateNewIndexFile(&conf, vindex)
	if err != nil {
		t.Error(err)
	}

	record, err := reader.FindByFilename(&conf, "test2.mp4")
	if err != nil {
		t.Error(err)
	}

	if record != nil {
		t.Error(errors.Errorf(""))
	}
}

func TestFindByDigest1(t *testing.T) {
	vindex := vindexdata.VIndex{
		vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4"),
	}
	err := writer.CreateNewIndexFile(&conf, vindex)
	if err != nil {
		t.Error(err)
	}

	record, err := reader.FindByDigest(&conf, vindex[0].Digest.Hex())
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(*record, vindex[0]); diff != "" {
		t.Error(diff)
	}
}

func TestFindByDigest2(t *testing.T) {
	vindex := vindexdata.VIndex{
		vindexdata.NewVIndexItem("video1", 20190101, "test1.mp4"),
	}
	err := writer.CreateNewIndexFile(&conf, vindex)
	if err != nil {
		t.Error(err)
	}

	record, err := reader.FindByDigest(&conf, "a4ddbc2cc9364c45b0012d9acdf32cf0")
	if err != nil {
		t.Error(err)
	}

	if record != nil {
		t.Error(errors.Errorf(""))
	}
}
