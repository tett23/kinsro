package main

import (
	"syscall/js"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func main() {
	vindex, err := LoadVIndex()
	if err != nil {
		panic(err)
	}

	js.Global().Set("vindex", vindex)
}

func LoadVIndex() (vindexdata.VIndex, error) {
	return vindexdata.VIndex{
		vindexdata.NewVIndexItem("test", 20181020, "test"),
	}, nil
}
