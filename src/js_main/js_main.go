package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func main() {
	vindex, err := LoadVIndex()
	if err != nil {
		panic(err)
	}

	data, err := json.Marshal(vindex)
	if err != nil {
		panic(err)
	}

	js.Global().Set("vindex", string(data))
}

// LoadVIndex LoadVIndex
func LoadVIndex() (vindexdata.VIndex, error) {
	return vindexdata.VIndex{
		vindexdata.NewVIndexItem("test", 20181020, "test"),
	}, nil
}
