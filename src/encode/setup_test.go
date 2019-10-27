package encode_test

import (
	"os"
	"testing"

	"github.com/tett23/kinsro/src/tests"
)

func TestMain(m *testing.M) {
	tests.Setup()

	code := m.Run()

	os.Exit(code)
}
