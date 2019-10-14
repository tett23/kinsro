package main

import (
	"fmt"

	"github.com/tett23/kinsro/src/commands"
	"github.com/tett23/kinsro/src/config"
	"github.com/tett23/kinsro/src/vindex/writer"
)

func main() {
	conf := config.GetConfig()
	storagePaths := conf.StoragePaths

	vindex, err := commands.BuildVIndex(storagePaths)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	err = writer.CreateNewIndexFile(conf, vindex)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}
}
