package tests

import (
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/filesystem"
)

// Setup Setup
func Setup() {
	filesystem.CopyDotfile()
	config.Initialize()
}
