package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/pkg/errors"
	"github.com/tett23/kinsro/src/js_main/assets"
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
	js.Global().Set("vf", js.FuncOf(vf))

	a := make(chan int)
	<-a
}

// LoadVIndex LoadVIndex
func LoadVIndex() (vindexdata.VIndex, error) {
	vindex, err := vindexdata.NewVIndexFromBinary(assets.AssetIndex)
	if err != nil {
		return nil, errors.Wrapf(err, "NewVIndexFromBinary failed")
	}

	return vindex, nil
}

func vf(this js.Value, args []js.Value) interface{} {
	fmt.Println("hogehoge")
	return "hogeあああ"
}
