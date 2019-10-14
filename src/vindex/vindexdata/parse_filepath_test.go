package vindexdata

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParseFilepath1(t *testing.T) {
	path := "/test/2019/10/10/test.mp4"
	storagePaths := []string{"/test"}
	actual, err := ParseFilepath(storagePaths, path)
	if err != nil {
		t.Error(err)
	}

	expected := VIndexItem{
		Storage:  "test",
		Date:     20191010,
		Filename: "/test/2019/10/10/test.mp4",
	}
	if diff := cmp.Diff(actual, &expected); diff != "" {
		t.Errorf(diff)
	}
}

func TestParseFilepath2(t *testing.T) {
	path := "/test/2019/10/10/test.mp4"
	storagePaths := []string{""}
	_, err := ParseFilepath(storagePaths, path)

	if err == nil {
		t.Fail()
	}
}

func TestParseFilepath3(t *testing.T) {
	path := "/test/2019/10/10/test.ts"
	storagePaths := []string{"/test"}
	_, err := ParseFilepath(storagePaths, path)

	if err == nil {
		t.Fail()
	}
}

func TestParseFilepath4(t *testing.T) {
	path := "/test/2019/10/10/test.ts"
	storagePaths := []string{"/hoge"}
	_, err := ParseFilepath(storagePaths, path)

	if err == nil {
		t.Fail()
	}
}
