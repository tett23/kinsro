package main

import (
	"encoding/json"
	"syscall/js"

	"github.com/tett23/kinsro/src/vindex/vindexdata"
)

func main() {
	ch := make(chan int)

	js.Global().Set("parseVIndex", js.FuncOf(ParseVIndex))

	<-ch
}

// ParseVIndex ParseVIndex
func ParseVIndex(this js.Value, args []js.Value) interface{} {
	data := make([]byte, args[0].Get("length").Int())
	js.CopyBytesToGo(data, args[0])

	vindex, err := vindexdata.NewVIndexFromBinary(data)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.Marshal(vindex)
	if err != nil {
		panic(err)
	}

	js.Global().Set("vindex", string(jsonData))

	return nil
}
